package logic

import (
	handler2 "call_center/db/rpc/internal/handler"
	"call_center/public/exception"
	"context"
	"io"
	"log"

	"call_center/db/rpc/internal/svc"
	"call_center/db/rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type DbLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDbLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DbLoginLogic {
	return &DbLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DbLoginLogic) DbLogin(stream pb.Db_DbLoginServer) error {
	var err error
	exception.Try(func() {
		log.Println("<DbLogin>")
		for {
			req, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
				}
				log.Println("<DbLogin> 连接断开, req:", req)
				break
			}

			cmd := req.Cmd
			cmdType := cmd.CmdType
			switch cmdType {
			case pb.EDbCommand_E_DB_COMMAND_PUSH_CHAT_LOG:
				handler2.PushChatLog(l.svcCtx, cmd)
				break
			case pb.EDbCommand_E_DB_COMMAND_PUSH_CHAT_RECORD:
				handler2.PushChatRecord(l.svcCtx, cmd)
				break
			}

			if err != nil {
				continue
			}
		}
	}).Catch(func(e exception.Exception) {
		err = e.(error)
	}).Finally(func() {

	})

	return err
}
