package svc

import (
	"call_center/call/rpc/internal/config"
	"call_center/call/rpc/internal/interaction"
	"call_center/db/rpc/db"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Db     interaction.InterDb
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbRpc := db.NewDb(zrpc.MustNewClient(c.DbRpc))
	interDb := interaction.NewInterDb(dbRpc)
	return &ServiceContext{
		Config: c,
		Db:     interDb,
	}
}
