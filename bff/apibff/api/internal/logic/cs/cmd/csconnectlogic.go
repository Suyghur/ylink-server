package cmd

import (
	"context"

	"ylink/bff/apibff/api/internal/svc"
	"ylink/bff/apibff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsConnectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCsConnectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsConnectLogic {
	return &CsConnectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CsConnectLogic) CsConnect(req *types.PlayerInfo) (resp *types.CommResp, err error) {
	// todo: add your logic here and delete this line

	return
}
