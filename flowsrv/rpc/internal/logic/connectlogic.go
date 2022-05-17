package logic

import (
	"context"
	"ylink/comm/result"
	"ylink/core/auth/rpc/auth"
	"ylink/flowsrv/rpc/internal/mgr"

	"ylink/flowsrv/rpc/internal/svc"
	"ylink/flowsrv/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConnectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConnectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConnectLogic {
	return &ConnectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConnectLogic) Connect(in *pb.CommandReq, stream pb.Flowsrv_ConnectServer) error {
	authResp, err := l.svcCtx.AuthRpc.CheckAuth(l.ctx, &auth.CheckAuthReq{
		Type:        in.Type,
		AccessToken: in.AccessToken,
	})
	if err != nil {
		return stream.Send(&pb.CommandResp{
			Code: result.TokenParseError,
			Msg:  err.Error(),
			Data: nil,
		})
	}
	// update(对接的user的状态也返回)
	//stream.RecvMsg()
	mgr.GetFlowMgrInstance().SetFlow(authResp.Uid, stream)

	return stream.Send(&pb.CommandResp{
		Code: result.Ok,
		Msg:  "success",
		Data: nil,
	})
}
