package logic

import (
	"context"
	"ylink/apis/auth/pb"

	"ylink/bff/authbff/internal/svc"
	"ylink/bff/authbff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCsAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsAuthLogic {
	return &CsAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CsAuthLogic) CsAuth(req *types.CsAuthReq) (resp *types.AuthResp, err error) {
	if authResp, err := l.svcCtx.AuthRpc.CsAuth(l.ctx, &pb.CsAuthReq{
		CsId: req.CsId,
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
