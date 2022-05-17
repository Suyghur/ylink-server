//@File     producer.go
//@Time     2022/05/06
//@Author   #Suyghur,

package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/zeromicro/go-zero/core/logx"
	"go.opentelemetry.io/otel/attribute"
	"ylink/comm/ctxdata"
	"ylink/comm/trace"
)

type Producer struct {
	topic    string
	addr     []string
	config   *sarama.Config
	producer sarama.SyncProducer
}

func NewKafkaProducer(addr []string, topic string) *Producer {
	logx.Infof("brokers: %v", addr)
	logx.Infof("topic: %s", topic)
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
		panic(err.Error())
	}
	p.producer = producer
	return &p
}

func (p *Producer) SendMessage(ctx context.Context, m string, key ...string) (partition int32, offset int64, err error) {
	traceId := ctxdata.GetTraceIdFromCtx(ctx)
	msg := &sarama.ProducerMessage{}
	msg.Headers = []sarama.RecordHeader{{
		Key:   sarama.ByteEncoder("trace_id"),
		Value: sarama.ByteEncoder(traceId),
	}}
	msg.Topic = p.topic
	if len(key) == 1 {
		msg.Key = sarama.StringEncoder(key[0])
	}
	msg.Value = sarama.StringEncoder(m)

	trace.StartTrace(ctx, "SendMessageToKafka", func(ctx context.Context) {
		partition, offset, err = p.producer.SendMessage(msg)
	}, attribute.StringSlice("keys", key), attribute.String("topic", p.topic))
	return
}
