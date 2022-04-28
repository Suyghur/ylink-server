package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"ylink/bff/authbff/api/internal/svc"
	"ylink/bff/authbff/api/internal/types"
	"ylink/core/auth/rpc/auth"
)

type CsLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCsLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsLoginLogic {
	return &CsLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CsLoginLogic) CsLogin(req *types.CsAuthReq) (resp *types.AuthResp, err error) {
	authResp, err := l.svcCtx.AuthRpc.CsAuth(l.ctx, &auth.CsAuthReq{
		CsId: req.CsId,
	})
	if err != nil {
		return nil, err
	}
	return &types.AuthResp{
		AccessToken: authResp.AccessToken,
	}, nil
}
