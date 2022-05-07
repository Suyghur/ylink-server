package svc

import (
	"ylink/core/cmd/rpc/internal/config"
	"ylink/ext/kafka"
)

type ServiceContext struct {
	Config          config.Config
	ChatMsgProducer *kafka.Producer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		ChatMsgProducer: kafka.NewKafkaProducer(c.KqChatMsgConf.Brokers, c.KqChatMsgConf.Topic),
	}
}
