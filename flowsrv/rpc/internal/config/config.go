package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"ylink/comm/kafka"
)

type Config struct {
	zrpc.RpcServerConf
	InnerRpcConf         zrpc.RpcClientConf
	KqMsgBoxConsumerConf kafka.KqConsumerConfig
	KqCmdBoxConsumerConf kafka.KqConsumerConfig
	JwtAuth              struct {
		AccessSecret string
		AccessExpire int64
	}
}
