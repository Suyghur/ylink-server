// Code generated by goctl. DO NOT EDIT!
// Source: inner.proto

package server

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	treemap "github.com/liyue201/gostl/ds/map"
	"github.com/zeromicro/go-zero/core/logx"
	"go.opentelemetry.io/otel/attribute"
	"ylink/comm/kafka"
	"ylink/comm/model"
	"ylink/comm/trace"
	"ylink/core/inner/rpc/internal/ext"

	"ylink/core/inner/rpc/internal/logic"
	"ylink/core/inner/rpc/internal/svc"
	"ylink/core/inner/rpc/pb"
)

type InnerServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedInnerServer
	ConsumerGroup *kafka.ConsumerGroup
}

func NewInnerServer(svcCtx *svc.ServiceContext) *InnerServer {
	server := &InnerServer{
		svcCtx: svcCtx,
		ConsumerGroup: kafka.NewConsumerGroup(&kafka.ConsumerGroupConfig{
			KafkaVersion:   sarama.V1_0_0_0,
			OffsetsInitial: sarama.OffsetNewest,
			IsReturnErr:    false,
		},
			svcCtx.Config.KqMsgBoxConsumerConf.Brokers,
			[]string{svcCtx.Config.KqMsgBoxConsumerConf.Topic},
			svcCtx.Config.KqMsgBoxConsumerConf.GroupId),
	}
	server.subscribe()
	return server

}

func (s *InnerServer) PlayerFetchCsInfo(ctx context.Context, in *pb.InnerPlayerFetchCsInfoReq) (*pb.InnerPlayerFetchCsInfoResp, error) {
	l := logic.NewPlayerFetchCsInfoLogic(ctx, s.svcCtx)
	return l.PlayerFetchCsInfo(in)
}

func (s *InnerServer) PlayerDisconnect(ctx context.Context, in *pb.InnerPlayerDisconnectReq) (*pb.InnerPlayerDisconnectResp, error) {
	l := logic.NewPlayerDisconnectLogic(ctx, s.svcCtx)
	return l.PlayerDisconnect(in)
}

func (s *InnerServer) CsFetchPlayerQueue(ctx context.Context, in *pb.InnerCsFetchPlayerQueueReq) (*pb.InnerCsFetchPlayerQueueResp, error) {
	l := logic.NewCsFetchPlayerQueueLogic(ctx, s.svcCtx)
	return l.CsFetchPlayerQueue(in)
}

func (s *InnerServer) CsConnectPlayer(ctx context.Context, in *pb.InnerCsConnectPlayerReq) (*pb.InnerCsConnectPlayerResp, error) {
	l := logic.NewCsConnectPlayerLogic(ctx, s.svcCtx)
	return l.CsConnectPlayer(in)
}

func (s *InnerServer) UpdateUserStatus(ctx context.Context, in *pb.UpdateUserStatusReq) (*pb.UpdateUserStatusResp, error) {
	l := logic.NewUpdateUserStatusLogic(ctx, s.svcCtx)
	return l.UpdateUserStatus(in)
}

func (s *InnerServer) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (s *InnerServer) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (s *InnerServer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		if msg.Topic == s.svcCtx.Config.KqMsgBoxConsumerConf.Topic {
			s.handleMessage(sess, msg)
		}
	}
	return nil
}

func (s *InnerServer) handleMessage(sess sarama.ConsumerGroupSession, msg *sarama.ConsumerMessage) {
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
		trace.StartTrace(ctx, "InnerServer.handleMessage.SendMessage", func(ctx context.Context) {
			if len(message.ReceiverId) == 0 || message.ReceiverId == "" {
				// 玩家发的消息
				p2cMap := ext.Game2PlayerMap.Get(message.GameId).(*treemap.Map)
				message.ReceiverId = p2cMap.Get(message.SenderId).(string)
				logx.WithContext(ctx).Infof("receiver: %s", message.ReceiverId)
				kMsg, _ := json.Marshal(message)
				s.svcCtx.KqMsgBoxProducer.SendMessage(ctx, string(kMsg), message.ReceiverId)
			} else {
				s.svcCtx.KqMsgBoxProducer.SendMessage(ctx, string(msg.Value), message.ReceiverId)
			}
			sess.MarkMessage(msg, "")
		}, attribute.String("msg.key", string(msg.Key)))
	})
}

func (s *InnerServer) subscribe() {
	go s.ConsumerGroup.RegisterHandleAndConsumer(s)
}
