package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"ylink/comm/kafka"
)

type Config struct {
	zrpc.RpcServerConf
	KqMsgBoxConsumerConf kafka.KqConsumerConfig
	KqMsgBoxProducerConf kafka.KqProducerConfig
	KqCmdBoxProducerConf kafka.KqProducerConfig
	//KqDbBoxProducerConf  kafka.KqProducerConfig

	MockFolder string
}
