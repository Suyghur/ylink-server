package logic

import (
	"call_center/call/rpc/internal/core"
	"call_center/call/rpc/internal/handler"
	"call_center/call/rpc/pb"
	"call_center/public/exception"
	"context"
	"fmt"
	"log"

	"call_center/call/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type ClientCallLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClientCallLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientCallLogic {
	return &ClientCallLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClientCallLogic) ClientCall(req *pb.ClientMsgReq) (*pb.ClientMsgRes, error) {
	var resList []*pb.CommandMsg
	var err error
	server := core.GetServer()

	exception.Try(func() {
		if req.IdInfo == nil {
			errStr := fmt.Sprintf("<ClientCall> IdInfo is nil")
			err = exception.MakeError(int32(pb.EErrorCode_ERR_PARAM_ERROR), errStr)
			exception.Throw(err)
		}

		id := req.IdInfo.Id
		stream := server.GetPlayerStream(id)
		if stream == nil {
			errStr := fmt.Sprintf("<ClientCall> client id not login, id:%v", id)
			err = exception.MakeError(int32(pb.EErrorCode_ERR_PARAM_ERROR), errStr)
			exception.Throw(errStr)
		}

		// 接收玩家信息
		for _, cmd := range req.Cmd {
			cmdType := cmd.CmdType
			cmdRes := new(pb.CommandMsg)
			switch cmdType {
			case pb.ECommand_CALL_PLAYER_MSG:
				// 玩家发消息
				err = handler.FromPlayerMsg(l.svcCtx.Db, server, id, cmd)
				break
			case pb.ECommand_CALL_PLAYER_LOGOUT:
				// 玩家退出
				handler.PlayerLogout(server, id)
				break
			default:
				errStr := fmt.Sprintf("<ClientCall> invalid cmd type:%s", cmdType)
				log.Println(errStr)
				exception.Throw(errStr)
				break
			}
			if err != nil {
				continue
			}
			resList = append(resList, cmdRes)
		}
	}).Catch(func(e exception.Exception) {
		log.Println("<ClientCall> error: ", e)
		err = e.(error)
	}).Finally(func() {
		// server.OnPlayerDisConnect(server, id)
	})
	return &pb.ClientMsgRes{Cmd: resList}, err
}
