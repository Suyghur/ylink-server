package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"ylink/ext/kafka"
)

type Config struct {
	zrpc.RpcServerConf
	KqChatMsgConf kafka.KqConfig
}
