package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"ylink/bff/apibff/api/internal/config"
	"ylink/gateway/rpc/gateway"
)

type ServiceContext struct {
	Config     config.Config
	GatewayRpc gateway.Gateway
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		GatewayRpc: gateway.NewGateway(zrpc.MustNewClient(c.GatewayRpcConf)),
	}
}
