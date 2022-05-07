//@File     kqconfig.go
//@Time     2022/05/06
//@Author   #Suyghur,

package kafka

type KqConfig struct {
	Brokers []string
	Topic   string
	GroupId string
}
