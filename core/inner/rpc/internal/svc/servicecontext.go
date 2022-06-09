package svc

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/bytedance/sonic"
	"github.com/gookit/event"
	"github.com/liyue201/gostl/ds/list/simplelist"
	treemap "github.com/liyue201/gostl/ds/map"
	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/logx"
	"go.opentelemetry.io/otel/attribute"
	"io/ioutil"
	"ylink/comm/globalkey"
	"ylink/comm/kafka"
	"ylink/comm/model"
	"ylink/comm/trace"
	"ylink/core/inner/rpc/internal/config"
	"ylink/core/inner/rpc/internal/ext"
)

type ServiceContext struct {
	Config             config.Config
	KqMsgBoxProducer   *kafka.Producer
	KqCmdBoxProducer   *kafka.Producer
	KqMsgConsumerGroup *kafka.ConsumerGroup
	TimeoutCron        *cron.Cron
}

func NewServiceContext(c config.Config) *ServiceContext {
	svcCtx := &ServiceContext{
		Config:           c,
		KqMsgBoxProducer: kafka.NewKafkaProducer(c.KqMsgBoxProducerConf.Brokers, c.KqMsgBoxProducerConf.Topic),
		KqCmdBoxProducer: kafka.NewKafkaProducer(c.KqCmdBoxProducerConf.Brokers, c.KqCmdBoxProducerConf.Topic),
		KqMsgConsumerGroup: kafka.NewConsumerGroup(&kafka.ConsumerGroupConfig{
			KafkaVersion:   sarama.V1_0_0_0,
			OffsetsInitial: sarama.OffsetNewest,
			IsReturnErr:    false,
		},
			c.KqMsgBoxConsumerConf.Brokers,
			[]string{c.KqMsgBoxConsumerConf.Topic},
			c.KqMsgBoxConsumerConf.GroupId),
	}
	go svcCtx.subscribe()
	fetchCsCenterInfo(c)
	return svcCtx
}

func (s *ServiceContext) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (s *ServiceContext) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (s *ServiceContext) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		if msg.Topic == s.Config.KqMsgBoxConsumerConf.Topic {
			s.handleMessage(sess, msg)
		}
	}
	return nil
}

func (s *ServiceContext) handleMessage(sess sarama.ConsumerGroupSession, msg *sarama.ConsumerMessage) {
	traceId := kafka.GetTraceFromHeader(msg.Headers)
	if len(traceId) == 0 {
		return
	}
	trace.RunOnTracing(traceId, func(ctx context.Context) {
		var message model.KqMessage
		if err := sonic.Unmarshal(msg.Value, &message); err != nil {
			logx.WithContext(ctx).Errorf("unmarshal msg error: %v", err)
			return
		}
		logx.WithContext(ctx).Infof("handle message: %s", msg.Value)
		trace.StartTrace(ctx, "InnerServer.handleMessage.SendMessage", func(ctx context.Context) {
			if len(message.ReceiverId) == 0 || message.ReceiverId == "" {
				// receiverId为空代表这条消息是玩家发送的
				// 玩家发的消息，先从connectedMap找对应的客服，没有则从vipMap找，都没有则丢弃信息不投递
				if playerInfo := ext.GetConnectedPlayerInfo(message.GameId, message.Uid); playerInfo != nil {
					message.ReceiverId = playerInfo.CsId
				} else {
					if playerInfo := ext.GetVipPlayer(message.GameId, message.Uid); playerInfo != nil {
						message.ReceiverId = playerInfo.CsId
					} else {
						message.ReceiverId = ""
					}
				}

				// 经过填补后receiver_id还是空的则有异常，丢弃信息不投递
				if len(message.ReceiverId) != 0 && message.ReceiverId != "" {
					logx.WithContext(ctx).Infof("receiver: %s", message.ReceiverId)
					kMsg, _ := sonic.MarshalString(message)
					s.KqMsgBoxProducer.SendMessage(ctx, kMsg, message.ReceiverId)
				} else {
					logx.WithContext(ctx).Errorf("can not find receiver of the sender")
				}
			} else {
				// receiverId不为空代表这条消息是客服发的
				s.KqMsgBoxProducer.SendMessage(ctx, string(msg.Value), message.ReceiverId)
			}
			sess.MarkMessage(msg, "")
		}, attribute.String("msg.key", string(msg.Key)))
	})
}

func (s *ServiceContext) subscribe() {
	go s.KqMsgConsumerGroup.RegisterHandleAndConsumer(s)

	// 注册事件
	event.On(globalkey.EventRemoveTimeoutJob, event.ListenerFunc(func(e event.Event) error {
		logx.Info("on event remove timeout job...")
		entryId := e.Get("entry_id").(cron.EntryID)

		s.TimeoutCron.Remove(entryId)
		return nil
	}), event.High)

	// 初始化定时任务
	s.TimeoutCron = cron.New(cron.WithSeconds())
	s.TimeoutCron.Start()
}

func fetchCsCenterInfo(c config.Config) {
	// mock info
	ext.CsInfoMap = treemap.New(treemap.WithGoroutineSafe())
	ext.GameVipMap = treemap.New(treemap.WithGoroutineSafe())
	ext.GameOnlinePlayerMap = treemap.New(treemap.WithGoroutineSafe())
	ext.GameConnectedMap = treemap.New(treemap.WithGoroutineSafe())
	ext.WaitingList = simplelist.New()
	go loadMockInfo(c)
}

func loadMockInfo(c config.Config) {
	// 加载游戏列表
	logx.Info("加载游戏列表")
	var gameIds []string
	gameIdsData, err := ioutil.ReadFile(c.MockFolder + "/game_id.json")
	if err != nil {
		logx.Errorf("parse game_id.json has some error: %v", err)
		return
	}
	if err := sonic.Unmarshal(gameIdsData, &gameIds); err != nil {
		return
	}

	// 加载vip玩家信息
	logx.Info("加载vip玩家信息")
	for _, gameId := range gameIds {
		vipPlayerMap := treemap.New(treemap.WithGoroutineSafe())
		var playerInfos []*model.PlayerInfo
		playerInfosData, _ := ioutil.ReadFile(c.MockFolder + "/" + gameId + ".json")
		if err := sonic.Unmarshal(playerInfosData, &playerInfos); err != nil {
			return
		}
		for _, playerInfo := range playerInfos {
			vipPlayerMap.Insert(playerInfo.PlayerId, playerInfo)
		}
		ext.GameVipMap.Insert(gameId, vipPlayerMap)
	}

	// 加载客服信息
	logx.Info("加载客服信息")
	var csInfos []*model.CsInfo
	csInfosData, err := ioutil.ReadFile(c.MockFolder + "/cs_info.json")
	if err := sonic.Unmarshal(csInfosData, &csInfos); err != nil {
		return
	}
	for _, csInfo := range csInfos {
		ext.CsInfoMap.Insert(csInfo.CsId, csInfo)
	}
}
