package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"ylink/core/inner/rpc/inner"
	"ylink/flowsrv/rpc/internal/config"
)

type ServiceContext struct {
	Config   config.Config
	InnerRpc inner.Inner
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		InnerRpc: inner.NewInner(zrpc.MustNewClient(c.InnerRpcConf)),
	}
}
