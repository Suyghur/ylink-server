package logic

import (
	handler2 "call_center/db/rpc/internal/handler"
	"call_center/public/exception"
	"context"
	"log"

	"call_center/db/rpc/internal/svc"
	"call_center/db/rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type DbCallLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDbCallLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DbCallLogic {
	return &DbCallLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DbCallLogic) DbCall(in *pb.DbMsgReq) (*pb.DbMsgRes, error) {
	var err error
	res := new(pb.DbMsgRes)
	exception.Try(func() {
		cmd := in.GetCmd()
		cmdType := cmd.CmdType
		switch cmdType {
		case pb.EDbCommand_E_DB_COMMAND_GET_CONFIG:
			res.Cmd = handler2.GetConfig(l.svcCtx)
			break
		case pb.EDbCommand_E_DB_COMMAND_GET_CHAT_RECORD:
			res.Cmd = handler2.GetChatRecord(l.svcCtx, cmd)
			break
		case pb.EDbCommand_E_DB_COMMAND_GET_CHAT_LOG:
			res.Cmd = handler2.GetChatLog(l.svcCtx, cmd)
			break
		}
		log.Println("<DbCall> cmdType: ", cmdType)
	}).Catch(func(e exception.Exception) {
		err = e.(error)
	}).Finally(func() {

	})
	return res, err
}
