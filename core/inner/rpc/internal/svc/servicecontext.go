package svc

import (
	treemap "github.com/liyue201/gostl/ds/map"
	"github.com/liyue201/gostl/ds/set"
	"ylink/comm/kafka"
	"ylink/comm/model"
	"ylink/core/inner/rpc/internal/config"
	"ylink/core/inner/rpc/internal/ext"
)

type ServiceContext struct {
	Config           config.Config
	KqMsgBoxProducer *kafka.Producer
}

func NewServiceContext(c config.Config) *ServiceContext {
	fetchCsCenterInfo()
	ext.Game2PlayerStatMap = treemap.New(treemap.WithGoroutineSafe())
	ext.CsStatSet = set.New(set.WithGoroutineSafe())
	return &ServiceContext{
		Config:           c,
		KqMsgBoxProducer: kafka.NewKafkaProducer(c.KqMsgBoxProducerConf.Brokers, c.KqMsgBoxProducerConf.Topic),
	}
}

func fetchCsCenterInfo() {
	// mock info
	mockInfo()
}

func mockInfo() {
	ext.Game2PlayerMap = treemap.New(treemap.WithGoroutineSafe())
	ext.CsMap = treemap.New(treemap.WithGoroutineSafe())

	// 已连接的映射

	// 专属客服映射
	game1231P2cMap := treemap.New(treemap.WithGoroutineSafe())
	game1231P2cMap.Insert("player1231", "cs_1231")
	game1231P2cMap.Insert("player1111", "cs_2222")

	game1111P2cMap := treemap.New(treemap.WithGoroutineSafe())
	game1111P2cMap.Insert("player1231", "cs_1111")

	ext.Game2PlayerMap.Insert("game1231", game1231P2cMap)
	ext.Game2PlayerMap.Insert("game1111", game1111P2cMap)

	ext.CsMap.Insert("cs_1231", &model.CsInfo{
		CsId:         "cs_1231",
		CsNickname:   "客服1231",
		CsAvatarUrl:  "https://www.baidu.com",
		CsSignature:  "我是客服1231",
		OnlineStatus: 1,
	})
	ext.CsMap.Insert("cs_1111", &model.CsInfo{
		CsId:         "cs_1111",
		CsNickname:   "客服1111",
		CsAvatarUrl:  "https://www.baidu.com",
		CsSignature:  "我是客服1111",
		OnlineStatus: 1,
	})
	ext.CsMap.Insert("cs_2222", &model.CsInfo{
		CsId:         "cs_2222",
		CsNickname:   "客服2222",
		CsAvatarUrl:  "https://www.baidu.com",
		CsSignature:  "我是客服2222",
		OnlineStatus: 0,
	})
}
