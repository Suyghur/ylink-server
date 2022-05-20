package svc

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/zeromicro/go-zero/core/logx"
	gozerotrace "github.com/zeromicro/go-zero/core/trace"
	"github.com/zeromicro/go-zero/zrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	oteltrace "go.opentelemetry.io/otel/trace"
	"net/http"
	"ylink/comm/kafka"
	"ylink/comm/model"
	"ylink/comm/trace"
	"ylink/comm/utils"
	"ylink/core/inner/rpc/inner"
	"ylink/flowsrv/rpc/internal/config"
)

type ServiceContext struct {
	Config        config.Config
	InnerRpc      inner.Inner
	ConsumerGroup *kafka.ConsumerGroup
}

func NewServiceContext(c config.Config) *ServiceContext {
	svcCtx := &ServiceContext{
		Config:   c,
		InnerRpc: inner.NewInner(zrpc.MustNewClient(c.InnerRpcConf)),
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

func (s *ServiceContext) runWithCtx(callback func(ctx context.Context), kv ...attribute.KeyValue) {
	propagator := otel.GetTextMapPropagator()
	tracer := otel.GetTracerProvider().Tracer(gozerotrace.TraceName)
	ctx := propagator.Extract(context.Background(), propagation.HeaderCarrier(http.Header{}))
	spanName := utils.CallerFuncName()
	spanCtx, span := tracer.Start(ctx, spanName, oteltrace.WithSpanKind(oteltrace.SpanKindServer), oteltrace.WithAttributes(kv...))
	defer span.End()
	propagator.Inject(spanCtx, propagation.HeaderCarrier(http.Header{}))
	callback(spanCtx)
}

func (s *ServiceContext) handleMessage(sess sarama.ConsumerGroupSession, msg *sarama.ConsumerMessage) {
	traceId := kafka.GetTraceFromHeader(msg.Headers)
	if len(traceId) == 0 {
		return
	}
	trace.RunOnTracing(traceId, func(ctx context.Context) {
		var message model.KqMessage
		if err := json.Unmarshal(msg.Value, &message); err != nil {
			logx.Errorf("unmarshal msg error: %v", err)
			return
		}
		trace.StartTrace(ctx, "FlowsrvServer.handleMessage.PushMessage", func(ctx context.Context) {
			//if len(message.ReceiverId) == 0 || message.ReceiverId == "" {
			//	// 玩家发的消息
			//	p2cMap := ext.IdMap.Get(message.GameId).(*treemap.Map)
			//	message.ReceiverId = p2cMap.Get(message.SenderId).(string)
			//	logx.Infof("receiver: %s", message.ReceiverId)
			//	b, _ := json.Marshal(message)
			//	s.svcCtx.KqMsgBoxProducer.SendMessage(ctx, string(b), message.ReceiverId)
			//} else {
			//	s.svcCtx.KqMsgBoxProducer.SendMessage(ctx, string(msg.Value), message.ReceiverId)
			//}
			logx.WithContext(ctx).Infof("headers: %v", msg.Headers)
			logx.WithContext(ctx).Infof("traceId: %s", msg.Headers[0].Value)
			logx.WithContext(ctx).Infof("flowsrv recv message: %v", string(msg.Value))
			sess.MarkMessage(msg, "")
		}, attribute.String("msg.key", string(msg.Key)))
	})
}

func (s *ServiceContext) subscribe() {
	go s.ConsumerGroup.RegisterHandleAndConsumer(s)
}
