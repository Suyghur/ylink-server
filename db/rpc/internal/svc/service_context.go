package svc

import (
	"call_center/db/model"
	"call_center/db/rpc/internal/config"
	es "call_center/public/es"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	ConfigModel   model.ChatConfigsModel
	SensWordModel model.ChatSensitiveWordsModel
	Es            es.EsMgrInterface
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		ConfigModel:   model.NewChatConfigsModel(conn),
		SensWordModel: model.NewChatSensitiveWordsModel(conn),
		Es:            es.New(c.EsConf),
	}
}
