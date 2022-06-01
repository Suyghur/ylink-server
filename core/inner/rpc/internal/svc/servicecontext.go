package svc

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/liyue201/gostl/ds/list/simplelist"
	treemap "github.com/liyue201/gostl/ds/map"
	"github.com/zeromicro/go-zero/core/logx"
	"go.opentelemetry.io/otel/attribute"
	"ylink/comm/kafka"
	"ylink/comm/model"
	"ylink/comm/trace"
	"ylink/core/inner/rpc/internal/config"
	"ylink/core/inner/rpc/internal/ext"
)

type ServiceContext struct {
	Config           config.Config
	KqMsgBoxProducer *kafka.Producer
	ConsumerGroup    *kafka.ConsumerGroup
}

func NewServiceContext(c config.Config) *ServiceContext {
	fetchCsCenterInfo()
	svcCtx := &ServiceContext{
		Config:           c,
		KqMsgBoxProducer: kafka.NewKafkaProducer(c.KqMsgBoxProducerConf.Brokers, c.KqMsgBoxProducerConf.Topic),
		ConsumerGroup: kafka.NewConsumerGroup(&kafka.ConsumerGroupConfig{
			KafkaVersion:   sarama.V1_0_0_0,
			OffsetsInitial: sarama.OffsetNewest,
			IsReturnErr:    false,
		},
			c.KqMsgBoxConsumerConf.Brokers,
			[]string{c.KqMsgBoxConsumerConf.Topic},
			c.KqMsgBoxConsumerConf.GroupId),
	}
	go svcCtx.subscribe()
	go fetchCsCenterInfo()
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
		if err := json.Unmarshal(msg.Value, &message); err != nil {
			logx.WithContext(ctx).Errorf("unmarshal msg error: %v", err)
			return
		}
		logx.WithContext(ctx).Infof("handle message: %s", msg.Value)
		trace.StartTrace(ctx, "InnerServer.handleMessage.SendMessage", func(ctx context.Context) {
			if len(message.ReceiverId) == 0 || message.ReceiverId == "" {
				// 玩家发的消息，先从connMap找对应的客服，没有则从vipMap找，都没有则丢弃信息不投递
				if ext.GameConnMap.Contains(message.GameId) {
					// 先从connMap找对应的客服映射
					if playerConnMap := ext.GameConnMap.Get(message.GameId).(*treemap.Map); playerConnMap.Contains(message.SenderId) {
						message.ReceiverId = playerConnMap.Get(message.SenderId).(string)
					} else {
						if ext.GameVipMap.Contains(message.GameId) {
							// 从vipMap里面找
							if playerVipMap := ext.GameVipMap.Get(message.GameId).(*treemap.Map); playerVipMap.Contains(message.SenderId) {
								message.ReceiverId = playerVipMap.Get(message.SenderId).(string)
							} else {
								message.ReceiverId = ""
							}
						} else {
							message.ReceiverId = ""
						}
					}
				} else {
					if ext.GameVipMap.Contains(message.GameId) {
						// 从vipMap里面找
						if playerVipMap := ext.GameVipMap.Get(message.GameId).(*treemap.Map); playerVipMap.Contains(message.SenderId) {
							message.ReceiverId = playerVipMap.Get(message.SenderId).(string)
						} else {
							message.ReceiverId = ""
						}
					} else {
						message.ReceiverId = ""
					}
				}

				// 经过填补后receiver_id还是空的则有异常，丢弃信息不投递
				if len(message.ReceiverId) != 0 && message.ReceiverId != "" {
					logx.WithContext(ctx).Infof("receiver: %s", message.ReceiverId)
					kMsg, _ := json.Marshal(message)
					s.KqMsgBoxProducer.SendMessage(ctx, string(kMsg), message.ReceiverId)
				} else {
					logx.WithContext(ctx).Errorf("can not find receiver of the sender")
				}
			} else {
				s.KqMsgBoxProducer.SendMessage(ctx, string(msg.Value), message.ReceiverId)
			}
			sess.MarkMessage(msg, "")
		}, attribute.String("msg.key", string(msg.Key)))
	})
}

func (s *ServiceContext) subscribe() {
	go s.ConsumerGroup.RegisterHandleAndConsumer(s)
}

func fetchCsCenterInfo() {
	// mock info
	ext.Game2PlayerStatusMap = treemap.New(treemap.WithGoroutineSafe())
	ext.GameConnMap = treemap.New(treemap.WithGoroutineSafe())
	ext.CsInfoMap = treemap.New(treemap.WithGoroutineSafe())
	ext.WaitingQueue = simplelist.New()
	mockInfo()
}
func mockInfo() {
	ext.GameVipMap = treemap.New(treemap.WithGoroutineSafe())
	ext.CsInfoMap = treemap.New(treemap.WithGoroutineSafe())

	// 已连接的映射

	// 专属客服映射
	game1231P2cMap := treemap.New(treemap.WithGoroutineSafe())
	game1231P2cMap.Insert("player_1231", "cs_1231")
	game1231P2cMap.Insert("player_1111", "cs_2222")

	game1111P2cMap := treemap.New(treemap.WithGoroutineSafe())
	game1111P2cMap.Insert("player_1231", "cs_1111")

	ext.GameVipMap.Insert("game_1231", game1231P2cMap)
	ext.GameVipMap.Insert("game_1111", game1111P2cMap)

	ext.CsInfoMap.Insert("cs_1231", &model.CsInfo{
		CsId:         "cs_1231",
		CsNickname:   "客服1231",
		CsAvatarUrl:  "https://www.baidu.com",
		CsSignature:  "我是客服1231",
		OnlineStatus: 0,
	})
	ext.CsInfoMap.Insert("cs_1111", &model.CsInfo{
		CsId:         "cs_1111",
		CsNickname:   "客服1111",
		CsAvatarUrl:  "https://www.baidu.com",
		CsSignature:  "我是客服1111",
		OnlineStatus: 0,
	})
	ext.CsInfoMap.Insert("cs_2222", &model.CsInfo{
		CsId:         "cs_2222",
		CsNickname:   "客服2222",
		CsAvatarUrl:  "https://www.baidu.com",
		CsSignature:  "我是客服2222",
		OnlineStatus: 0,
	})
}
