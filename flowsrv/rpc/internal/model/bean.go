//@File     bean.go
//@Time     2022/6/1
//@Author   #Suyghur,

package model

import (
	"github.com/zeromicro/go-zero/core/logx"
	"ylink/flowsrv/rpc/internal/svc"
	"ylink/flowsrv/rpc/pb"
)

type Flow struct {
	EndFlow chan int
	Message chan string
	SvcCtx  *svc.ServiceContext
	Logger  logx.Logger
	Stream  pb.Flowsrv_ConnectServer
	Type    int32
	Uid     string
	GameId  string
	FlowId  string
}
