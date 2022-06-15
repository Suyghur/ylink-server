//@File     bean.go
//@Time     2022/6/1
//@Author   #Suyghur,

package model

import (
	"github.com/go-redis/redis/v8"
	"ylink/core/inner/rpc/inner"
	"ylink/flowsrv/rpc/pb"
)

type Flow struct {
	EndFlow     chan int
	Message     chan string
	Stream      pb.Flowsrv_ConnectServer
	RedisClient *redis.Client
	InnerRpc    inner.Inner
	Type        int32
	Uid         string
	GameId      string
	FlowId      string
}
