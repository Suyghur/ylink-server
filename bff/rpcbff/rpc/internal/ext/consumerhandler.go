//@File     consumerhandler.go
//@Time     2022/05/07
//@Author   #Suyghur,

package ext

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/zeromicro/go-zero/core/logx"
	"ylink/ext/kafka"
)

type callback func(msg []byte)

type ConsumerHandler struct {
	Callbacks     map[string]callback
	ConsumerGroup *kafka.ConsumerGroup
}

func (handler *ConsumerHandler) Init(config kafka.KqConfig) {
	handler.Callbacks = make(map[string]callback)
	handler.Callbacks[config.Topic] = handler.handleMessage

	consumerGroupConfig := kafka.ConsumerGroupConfig{
		KafkaVersion:   sarama.V2_8_0_0,
		OffsetsInitial: sarama.OffsetNewest,
		IsReturnErr:    false,
	}
	logx.WithContext(context.Background()).Infof("brokers: %v", config.Brokers)
	logx.WithContext(context.Background()).Infof("group id: %s", config.GroupId)
	handler.ConsumerGroup = kafka.NewConsumerGroup(&consumerGroupConfig, config.Brokers, []string{config.Topic}, config.GroupId)
}

func (handler *ConsumerHandler) handleMessage(msg []byte) {
	logx.WithContext(context.Background()).Infof("handle message from kafka: %s", string(msg))
	//msgFromMq:=
}

func (ConsumerHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (ConsumerHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (handler *ConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		logx.WithContext(context.Background()).Infof("kafka get info to mysql, topic: %s, partition: %d, msg: %s", msg.Topic, msg.Partition, string(msg.Value))
		handler.Callbacks[msg.Topic](msg.Value)
	}
	return nil
}
