package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"ylink/core/auth/rpc/auth"
	"ylink/flowsrv/rpc/internal/config"
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
