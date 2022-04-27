package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"ylink/apis/cmd/cmd"
	"ylink/bff/apibff/internal/config"
	"ylink/bff/apibff/internal/middleware"
	"ylink/ext/globalkey"
)

type ServiceContext struct {
	Config     config.Config
	CmdRpc     cmd.Cmd
	Player2Ctx rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	globalkey.AccessSecret = c.JwtAuth.AccessSecret
	globalkey.AccessExpire = c.JwtAuth.AccessExpire
	return &ServiceContext{
		Config:     c,
		CmdRpc:     cmd.NewCmd(zrpc.MustNewClient(c.CmdRpc)),
		Player2Ctx: middleware.NewPlayer2CtxMiddleware().Handle,
	}
}
