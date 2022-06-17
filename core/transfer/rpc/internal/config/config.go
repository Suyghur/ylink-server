package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"ylink/comm/es"
	"ylink/comm/kafka"
)

type Config struct {
	zrpc.RpcServerConf
	EsConf               es.EsConf
	KqMsgBoxConsumerConf kafka.KqConsumerConfig
}
