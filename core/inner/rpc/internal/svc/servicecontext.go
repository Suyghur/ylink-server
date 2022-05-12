package svc

import (
	"ylink/core/inner/rpc/internal/config"
	"ylink/core/inner/rpc/internal/ext"
	"ylink/ext/ds/treemap"
	"ylink/ext/kafka"
	"ylink/ext/model"
)

type ServiceContext struct {
	Config          config.Config
	RecvBoxProducer *kafka.Producer
}

func NewServiceContext(c config.Config) *ServiceContext {
	fetchCsCenterInfo()
	recvBoxProducer := kafka.NewKafkaProducer(c.KqSendMsgConf.Brokers, c.KqSendMsgConf.Topic)
	var sendBoxHandler ext.SendBoxConsumerHandler
	sendBoxHandler.Init(c.KqRecvMsgConf, recvBoxProducer)
	go sendBoxHandler.ConsumerGroup.RegisterHandleAndConsumer(&sendBoxHandler)
	return &ServiceContext{
		Config: c,
	}
}

func fetchCsCenterInfo() {
	// mock info
	mockInfo()
}

func mockInfo() {
	ext.IdMap = treemap.New(treemap.WithGoroutineSafe())
	ext.CsMap = treemap.New(treemap.WithGoroutineSafe())

	game1231P2cMap := treemap.New(treemap.WithGoroutineSafe())
	game1231P2cMap.Insert("player1231", "cs_1231")
	game1231P2cMap.Insert("player1111", "cs_2222")

	game1111P2cMap := treemap.New(treemap.WithGoroutineSafe())
	game1111P2cMap.Insert("player1231", "cs_1111")

	ext.IdMap.Insert("game1231", game1231P2cMap)
	ext.IdMap.Insert("game1111", game1111P2cMap)

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