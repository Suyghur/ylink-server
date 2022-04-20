package cmd

import (
	"context"

	"ylink/bff/apibff/api/internal/svc"
	"ylink/bff/apibff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsFetchPlayerInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCsFetchPlayerInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsFetchPlayerInfoLogic {
	return &CsFetchPlayerInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CsFetchPlayerInfoLogic) CsFetchPlayerInfo(req *types.CsFetchPlayerInfoReq) (resp *types.CommResp, err error) {
	// todo: add your logic here and delete this line

	return
}
