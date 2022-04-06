package cmd

import (
	"context"

	"ylink/bff/apibff/api/internal/svc"
	"ylink/bff/apibff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsLogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCsLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsLogoutLogic {
	return &CsLogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CsLogoutLogic) CsLogout() (resp *types.CommResp, err error) {
	// todo: add your logic here and delete this line

	return
}
