package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"ylink/bff/cmdbff/api/internal/config"
	"ylink/core/cmd/rpc/cmd"
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
