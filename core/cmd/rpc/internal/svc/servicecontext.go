package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"ylink/core/cmd/rpc/internal/config"
	"ylink/ext/kafka"
)

type ServiceContext struct {
	Config          config.Config
	ChatMsgProducer *kafka.Producer
	RedisClient     *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		ChatMsgProducer: kafka.NewKafkaProducer(c.KqChatMsgConf.Brokers, c.KqChatMsgConf.Topic),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
	}
}
