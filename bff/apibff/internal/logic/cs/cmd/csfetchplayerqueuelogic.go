package cmd

import (
	"context"

	"ylink/bff/apibff/internal/svc"
	"ylink/bff/apibff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsFetchPlayerQueueLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCsFetchPlayerQueueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsFetchPlayerQueueLogic {
	return &CsFetchPlayerQueueLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CsFetchPlayerQueueLogic) CsFetchPlayerQueue(req *types.CsFetchPlayerQueueReq) (resp *types.CommResp, err error) {
	// todo: add your logic here and delete this line

	return
}
