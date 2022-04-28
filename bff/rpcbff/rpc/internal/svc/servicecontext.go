package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"ylink/bff/rpcbff/rpc/internal/config"
	"ylink/core/auth/rpc/auth"
)

type ServiceContext struct {
	Config  config.Config
	AuthRpc auth.Auth
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		AuthRpc: auth.NewAuth(zrpc.MustNewClient(c.AuthRpcConf)),
	}
}
