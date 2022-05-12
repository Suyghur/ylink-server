package logic

import (
	"context"
	"ylink/core/auth/rpc/auth"
	"ylink/ext/result"

	"ylink/flowsrv/rpc/internal/svc"
	"ylink/flowsrv/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DisconnectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDisconnectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DisconnectLogic {
	return &DisconnectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DisconnectLogic) Disconnect(in *pb.CommandReq) (*pb.CommandResp, error) {
	_, err := l.svcCtx.AuthRpc.CheckAuth(l.ctx, &auth.CheckAuthReq{
		AccessToken: in.AccessToken,
	})
	//data, _ := structpb.NewStruct(treemap[string]interface{}{})
	if err != nil {
		return &pb.CommandResp{
			Code: result.TokenParseError,
			Msg:  err.Error(),
			Data: nil,
		}, err
	}
	return &pb.CommandResp{
		Code: result.Ok,
		Msg:  "success",
		Data: nil,
	}, nil
}
