package logic

import (
	"context"
	"ylink/apis/auth/pb"

	"ylink/bff/authbff/internal/svc"
	"ylink/bff/authbff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckAuthLogic {
	return &CheckAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckAuthLogic) CheckAuth(req *types.CheckAuthReq) (resp *types.AuthResp, err error) {
	if authResp, err := l.svcCtx.AuthRpc.CheckAuth(l.ctx, &pb.CheckAuthReq{
		Token: req.Token,
	}); err != nil {
		return &types.AuthResp{
			Code: authResp.Code,
			Msg:  authResp.Msg,
			Data: map[string]interface{}{},
		}, err
	} else {
		return &types.AuthResp{
			Code: authResp.Code,
			Msg:  authResp.Msg,
			Data: authResp.Data,
		}, nil
	}
}
