package config

import (
	es "call_center/public/es"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string
	}
	EsConf es.EsConfig
}
