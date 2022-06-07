//@File     bean.go
//@Time     2022/6/1
//@Author   #Suyghur,

package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"ylink/flowsrv/rpc/internal/svc"
	"ylink/flowsrv/rpc/pb"
)

type Flow struct {
	EndFlow chan int
	Message chan string
	Ctx     context.Context
	SvcCtx  *svc.ServiceContext
	Logger  logx.Logger
	Stream  pb.Flowsrv_ConnectServer
	Type    int64
	Uid     string
	GameId  string
	FlowId  string
}
