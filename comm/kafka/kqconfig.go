//@File     kqconfig.go
//@Time     2022/05/06
//@Author   #Suyghur,

package kafka

type KqProducerConfig struct {
	Brokers []string
	Topic   string
}

type KqConsumerConfig struct {
	Brokers []string
	Topic   string
	GroupId string
}
