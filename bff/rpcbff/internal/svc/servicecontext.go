package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"ylink/apis/auth/auth"
	"ylink/bff/rpcbff/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	AuthRpc auth.Auth
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		AuthRpc: auth.NewAuth(zrpc.MustNewClient(c.AuthRpc)),
	}
}
