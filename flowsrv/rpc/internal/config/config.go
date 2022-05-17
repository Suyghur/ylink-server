package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"ylink/comm/kafka"
)

type Config struct {
	zrpc.RpcServerConf
	AuthRpcConf          zrpc.RpcClientConf
	KqMsgBoxConsumerConf kafka.KqConsumerConfig
}
