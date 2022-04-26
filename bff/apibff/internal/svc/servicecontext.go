package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"ylink/apis/cmd/cmd"
	"ylink/bff/apibff/internal/config"
)

type ServiceContext struct {
	Config config.Config
	CmdRpc cmd.Cmd
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		CmdRpc: cmd.NewCmd(zrpc.MustNewClient(c.CmdRpc)),
	}
}
