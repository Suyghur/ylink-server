//@File     producer.go
//@Time     2022/05/06
//@Author   #Suyghur,

package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/zeromicro/go-zero/core/logx"
)

type Producer struct {
	topic    string
	addr     []string
	config   *sarama.Config
	producer sarama.SyncProducer
}

func NewKafkaProducer(addr []string, topic string) *Producer {
	p := Producer{}
	p.config = sarama.NewConfig()
	// Whether to enable the successes channel to be notified after the message is sent successfully
	p.config.Producer.Return.Successes = true
	// Set producer Message Reply level 0 1 all
	p.config.Producer.RequiredAcks = sarama.WaitForAll
	// Set the hash-key automatic hash partition. When sending a message, you must specify the key value of the message,
	// If there is no key, the partition will be selected randomly
	p.config.Producer.Partitioner = sarama.NewHashPartitioner

	p.addr = addr
	p.topic = topic

	// Initialize the client
	producer, err := sarama.NewSyncProducer(p.addr, p.config)
	if err != nil {
		logx.WithContext(context.Background()).Error(err.Error())
		return nil
	}
	p.producer = producer
	return &p
}

//func (p *Producer) SendMessage(m proto.Message, key ...string) (int32, int64, error) {
//	kMsg := &sarama.ProducerMessage{}
//	kMsg.Topic = p.topic
//	if len(key) == 1 {
//		kMsg.Key = sarama.StringEncoder(key[0])
//	}
//	bMsg, err := proto.Marshal(m)
//	if err != nil {
//		logx.WithContext(context.Background()).Errorf("proto marshal err: %s", err.Error())
//		return -1, -1, err
//	}
//	kMsg.Value = sarama.ByteEncoder(bMsg)
//	return p.producer.SendMessage(kMsg)
//}

func (p *Producer) SendMessage(m string, key ...string) (int32, int64, error) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = p.topic
	if len(key) == 1 {
		msg.Key = sarama.StringEncoder(key[0])
	}
	//bMsg, err := proto.Marshal(m)
	//if err != nil {
	//	logx.Errorf("proto marshal err: %s", err.Error())
	//	return -1, -1, err
	//}
	msg.Value = sarama.StringEncoder(m)
	return p.producer.SendMessage(msg)
}
