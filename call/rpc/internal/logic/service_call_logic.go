package logic

import (
	"call_center/call/rpc/internal/core"
	"call_center/call/rpc/internal/handler"
	"call_center/call/rpc/pb"
	"call_center/public/exception"
	"context"
	"errors"
	"fmt"
	"log"

	"call_center/call/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type ServiceCallLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewServiceCallLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ServiceCallLogic {
	return &ServiceCallLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ServiceCallLogic) ServiceCall(req *pb.ServiceMsgReq) (*pb.ServiceMsgRes, error) {
	var err error
	server := core.GetServer()
	id := req.IdInfo.Id
	var resList []*pb.CommandMsg
	exception.Try(func() {
		stream := server.GetServiceStream(id)
		if stream == nil {
			errStr := fmt.Sprintf("<ServiceCall> service id not login, id:%v", id)
			err = exception.MakeError(int32(pb.EErrorCode_ERR_PARAM_ERROR), errStr)
			exception.Throw(errStr)
		}
		// 接收客服信息
		for _, cmd := range req.Cmd {
			cmdType := cmd.CmdType
			cmdRes := new(pb.CommandMsg)
			switch cmdType {
			case pb.ECommand_CALL_SERVICE_MSG:
				err = handler.FromServiceMsg(server, id, cmd, l.svcCtx.Db)
				break
			case pb.ECommand_CALL_SERVICE_REPLY:
				err = handler.ServiceReply(server, id, cmd, l.svcCtx.Db)
				break
			case pb.ECommand_CALL_PLAYER_CHAT_LOG:
				cmdRes, err = handler.GetPlayerChatLog(id, cmd, l.svcCtx.Db)
				break
			default:
				err = errors.New(fmt.Sprintf("<ServiceCall> invalid cmd type:%s", cmdType))
				break
			}

			if err != nil {
				// 过滤
				exception.Throw(err)
				continue
			}
			resList = append(resList, cmdRes)
		}
	}).Catch(func(e exception.Exception) {
		log.Println("<ServiceCall> error: ", e)
		err = e.(error)
	}).Finally(func() {
		// server.OnServiceDisConnect(server, id)
	})
	return &pb.ServiceMsgRes{Cmd: resList}, err
}
