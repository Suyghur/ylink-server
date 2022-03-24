package config

import (
	"call_center/call/rpc/internal/core"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DbRpc    zrpc.RpcClientConf
	CoreConf core.Config
}
