//@File     sendboxhandler.go
//@Time     2022/05/12
//@Author   #Suyghur,

package ext

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/zeromicro/go-zero/core/logx"
	"ylink/ext/ds/treemap"
	"ylink/ext/kafka"
	"ylink/ext/model"
)

type callback func(msg []byte)

type SendBoxConsumerHandler struct {
	callbacks     map[string]callback
	producer      *kafka.Producer
	ConsumerGroup *kafka.ConsumerGroup
}

func (handler *SendBoxConsumerHandler) Init(c kafka.KqConfig, producer *kafka.Producer) {
	handler.callbacks = make(map[string]callback)
	handler.callbacks[c.Topic] = handler.handleMessage
	handler.producer = producer
	handler.ConsumerGroup = kafka.NewConsumerGroup(&kafka.ConsumerGroupConfig{
		KafkaVersion:   sarama.V0_10_2_0,
		OffsetsInitial: sarama.OffsetNewest,
		IsReturnErr:    false,
	}, c.Brokers, []string{c.Topic}, c.GroupId)
}

func (handler *SendBoxConsumerHandler) handleMessage(msg []byte) {
	logx.WithContext(context.Background()).Infof("message recv from send-box, %s", string(msg))
	// todo 将message转发到recv-box

	var message model.ChatMessage
	err := json.Unmarshal(msg, &message)
	if err != nil {
		logx.WithContext(context.Background()).Errorf("unmarshal message err: %s", err.Error())
		return
	}
	if len(message.ReceiverId) == 0 {
		// 玩家发的消息
		p2cMap := IdMap.Get(message.GameId).(*treemap.Map)
		message.ReceiverId = p2cMap.Get(message.SenderId).(string)
		b, _ := json.Marshal(message)
		handler.producer.SendMessage(string(b), message.ReceiverId)
	} else {
		handler.producer.SendMessage(string(msg), message.ReceiverId)
	}
}

func (SendBoxConsumerHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (SendBoxConsumerHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (handler *SendBoxConsumerHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		logx.WithContext(context.Background()).Infof("send-box get info to db, topic: %s, partition: %d, msg: %s", msg.Topic, msg.Partition, string(msg.Value))
		handler.callbacks[msg.Topic](msg.Value)
	}
	return nil
}
