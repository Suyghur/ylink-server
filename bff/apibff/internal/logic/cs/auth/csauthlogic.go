package auth

import (
	"context"
	"ylink/apis/auth/pb"

	"ylink/bff/apibff/internal/svc"
	"ylink/bff/apibff/internal/types"

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

func (l *CsAuthLogic) CsAuth(req *types.CsAuthReq) (resp *types.CommResp, err error) {
	if authResp, err := l.svcCtx.AuthRpc.CsAuth(l.ctx, &pb.CsAuthReq{
		Uname:    req.Uname,
		Password: req.Password,
	}); err != nil {
		return &types.CommResp{
			Code: authResp.Code,
			Msg:  "success",
			Data: map[string]interface{}{},
		}, err
	} else {
		return &types.CommResp{
			Code: 0,
			Msg:  "success",
			Data: authResp.Data,
		}, nil
	}
}
