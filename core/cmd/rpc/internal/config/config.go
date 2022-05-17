package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"ylink/comm/kafka"
)

type Config struct {
	zrpc.RpcServerConf
	InnerRpcConf         zrpc.RpcClientConf
	KqMsgBoxProducerConf kafka.KqProducerConfig
}
