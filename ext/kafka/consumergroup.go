//@File     consumergroup.go
//@Time     2022/05/06
//@Author   #Suyghur,

package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/zeromicro/go-zero/core/logx"
)

type ConsumerGroup struct {
	sarama.ConsumerGroup
	groupId string
	topics  []string
}

type ConsumerGroupConfig struct {
	KafkaVersion   sarama.KafkaVersion
	OffsetsInitial int64
	IsReturnErr    bool
}

func NewConsumerGroup(c *ConsumerGroupConfig, addr, topics []string, groupId string) *ConsumerGroup {
	config := sarama.NewConfig()
	config.Version = c.KafkaVersion
	config.Consumer.Offsets.Initial = c.OffsetsInitial
	config.Consumer.Return.Errors = c.IsReturnErr
	client, err := sarama.NewClient(addr, config)
	if err != nil {
		logx.WithContext(context.Background()).Error(err.Error())
		return nil
	}
	consumerGroup, err := sarama.NewConsumerGroupFromClient(groupId, client)
	if err != nil {
		logx.WithContext(context.Background()).Error(err.Error())
		return nil
	}
	return &ConsumerGroup{consumerGroup, groupId, topics}
}

func (cg *ConsumerGroup) RegisterHandleAndConsumer(handler sarama.ConsumerGroupHandler) {
	ctx := context.Background()
	for {
		if err := cg.ConsumerGroup.Consume(ctx, cg.topics, handler); err != nil {
			logx.Error(err.Error())
		}
	}
}
