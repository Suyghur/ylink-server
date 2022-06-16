package svc

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/bytedance/sonic"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	gozerotrace "github.com/zeromicro/go-zero/core/trace"
	"github.com/zeromicro/go-zero/zrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	oteltrace "go.opentelemetry.io/otel/trace"
	"net/http"
	"ylink/comm/globalkey"
	"ylink/comm/kafka"
	"ylink/comm/model"
	"ylink/comm/result"
	"ylink/comm/trace"
	"ylink/comm/utils"
	"ylink/core/inner/rpc/inner"
	"ylink/flowsrv/rpc/flowsrv"
	"ylink/flowsrv/rpc/internal/config"
	"ylink/flowsrv/rpc/internal/mgr"
	model2 "ylink/flowsrv/rpc/internal/model"
)

type ServiceContext struct {
	Config               config.Config
	InnerRpc             inner.Inner
	MessageConsumerGroup *kafka.ConsumerGroup
	CommandConsumerGroup *kafka.ConsumerGroup
	RedisClient          *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	svcCtx := &ServiceContext{
		Config:   c,
		InnerRpc: inner.NewInner(zrpc.MustNewClient(c.InnerRpcConf)),
		RedisClient: redis.NewClient(&redis.Options{
			Addr:     c.Redis.Host,
			Password: c.Redis.Pass,
		}),
		MessageConsumerGroup: kafka.NewConsumerGroup(&kafka.ConsumerGroupConfig{
			KafkaVersion:   sarama.V1_0_0_0,
			OffsetsInitial: sarama.OffsetNewest,
			IsReturnErr:    false,
		},
			c.KqMsgBoxConsumerConf.Brokers,
			[]string{c.KqMsgBoxConsumerConf.Topic},
			c.KqMsgBoxConsumerConf.GroupId),

		CommandConsumerGroup: kafka.NewConsumerGroup(&kafka.ConsumerGroupConfig{
			KafkaVersion:   sarama.V1_0_0_0,
			OffsetsInitial: sarama.OffsetNewest,
			IsReturnErr:    false,
		},
			c.KqCmdBoxConsumerConf.Brokers,
			[]string{c.KqCmdBoxConsumerConf.Topic},
			c.KqCmdBoxConsumerConf.GroupId),
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
		} else if msg.Topic == s.Config.KqCmdBoxConsumerConf.Topic {
			s.handleCommand(sess, msg)
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
		logx.WithContext(ctx).Infof("handle message: %s", msg.Value)
		var message model.KqMessage
		if err := sonic.Unmarshal(msg.Value, &message); err != nil {
			logx.Errorf("unmarshal msg error: %v", err)
			return
		}

		if message.Opt != model.CMD_SEND_MESSAGE {
			// 指令异常
			return
		}

		trace.StartTrace(ctx, "FlowsrvServer.handleMessage.PushMessage", func(ctx context.Context) {
			// 投递到receiver_id对应的redis队列暂存
			intCmd := s.RedisClient.LPush(ctx, message.ReceiverId, string(msg.Value))
			if size, err := intCmd.Result(); err != nil {
				logx.WithContext(ctx).Errorf("push message rmq err %v", err)
			} else {
				logx.WithContext(ctx).Infof("current rmq size: %d", size)
			}
			sess.MarkMessage(msg, "")
		}, attribute.String("msg.key", string(msg.Key)))
	})
}

func (s *ServiceContext) handleCommand(sess sarama.ConsumerGroupSession, msg *sarama.ConsumerMessage) {
	traceId := kafka.GetTraceFromHeader(msg.Headers)
	if len(traceId) == 0 {
		return
	}
	trace.RunOnTracing(traceId, func(ctx context.Context) {
		logx.WithContext(ctx).Infof("handle command: %s", msg.Value)
		var message model.KqMessage
		if err := sonic.Unmarshal(msg.Value, &message); err != nil {
			logx.Errorf("unmarshal msg error: %v", err)
			return
		}

		if message.Opt == model.CMD_SEND_MESSAGE {
			// 指令异常
			return
		}

		// 投递到receiver_id对应的redis队列暂存
		switch message.ReceiverId {
		case globalkey.All:
		case globalkey.AllPlayer:
		case globalkey.AllVipPlayer:
		case globalkey.AllNormalPlayer:
			for iter := mgr.GetFlowMgrInstance().All().Begin(); iter.IsValid(); iter.Next() {
				flow := iter.Value().(*model2.Flow)
				if flow.Type != globalkey.ConnectTypeNormalPlayer {
					continue
				}
				err := flow.Stream.Send(&flowsrv.CommandResp{
					Code: result.Ok,
					Msg:  "success",
					Data: msg.Value,
				})
				if err != nil {
					logx.WithContext(ctx).Errorf("%v", err)
				}
			}
		case globalkey.AllCs:
		default:
			trace.StartTrace(ctx, "FlowsrvServer.handleCommand.PushCmdMessage", func(ctx context.Context) {
				intCmd := s.RedisClient.LPush(ctx, message.ReceiverId, string(msg.Value))
				if size, err := intCmd.Result(); err != nil {
					logx.WithContext(ctx).Errorf("push message rmq err %v", err)
				} else {
					logx.WithContext(ctx).Infof("current rmq size: %d", size)
				}

			}, attribute.String("msg.key", string(msg.Key)))
			sess.MarkMessage(msg, "")
		}
	})
}

func (s *ServiceContext) subscribe() {
	go s.MessageConsumerGroup.RegisterHandleAndConsumer(s)
	go s.CommandConsumerGroup.RegisterHandleAndConsumer(s)
}
