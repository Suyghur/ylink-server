package cmd

import (
	"context"

	"ylink/bff/apibff/api/internal/svc"
	"ylink/bff/apibff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsDisconnectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCsDisconnectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsDisconnectLogic {
	return &CsDisconnectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CsDisconnectLogic) CsDisconnect() (resp *types.CommResp, err error) {
	// todo: add your logic here and delete this line

	return
}
