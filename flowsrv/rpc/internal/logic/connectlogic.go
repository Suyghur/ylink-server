package logic

import (
	"context"
	"ylink/core/auth/rpc/auth"
	"ylink/ext/result"

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
	_, err := l.svcCtx.AuthRpc.CheckAuth(l.ctx, &auth.CheckAuthReq{
		AccessToken: in.AccessToken,
	})
	//data, _ := structpb.NewStruct(map[string]interface{}{})
	if err != nil {
		return stream.Send(&pb.CommandResp{
			Code: result.TokenParseError,
			Msg:  err.Error(),
			Data: nil,
		})
	}
	return stream.Send(&pb.CommandResp{
		Code: result.Ok,
		Msg:  "success",
		Data: nil,
	})
}
