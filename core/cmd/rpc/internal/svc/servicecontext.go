package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"ylink/core/cmd/rpc/internal/config"
	"ylink/core/inner/rpc/inner"
	"ylink/ext/kafka"
)

type ServiceContext struct {
	Config          config.Config
	InnerRpc        inner.Inner
	SendBoxProducer *kafka.Producer
	RedisClient     *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		InnerRpc:        inner.NewInner(zrpc.MustNewClient(c.InnerRpcConf)),
		SendBoxProducer: kafka.NewKafkaProducer(c.KqSendMsgConf.Brokers, c.KqSendMsgConf.Topic),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
	}
}
