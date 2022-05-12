package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"ylink/ext/kafka"
)

type Config struct {
	zrpc.RpcServerConf
	KqRecvMsgConf kafka.KqConfig
	KqSendMsgConf kafka.KqConfig
}
