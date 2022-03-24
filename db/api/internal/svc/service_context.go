package svc

import (
	"call_center/db/api/internal/config"
	"call_center/db/rpc/db"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	DbRpc  db.Db
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DbRpc:  db.NewDb(zrpc.MustNewClient(c.DbRpc)),
	}
}
