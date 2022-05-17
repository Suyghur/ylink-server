package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"ylink/comm/kafka"
	"ylink/core/cmd/rpc/internal/config"
	"ylink/core/inner/rpc/inner"
)

type ServiceContext struct {
	Config           config.Config
	InnerRpc         inner.Inner
	KqMsgBoxProducer *kafka.Producer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		InnerRpc:         inner.NewInner(zrpc.MustNewClient(c.InnerRpcConf)),
		KqMsgBoxProducer: kafka.NewKafkaProducer(c.KqMsgBoxProducerConf.Brokers, c.KqMsgBoxProducerConf.Topic),
	}
}
