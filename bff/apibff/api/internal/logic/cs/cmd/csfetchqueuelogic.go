package cmd

import (
	"context"

	"ylink/bff/apibff/api/internal/svc"
	"ylink/bff/apibff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsFetchQueueLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCsFetchQueueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsFetchQueueLogic {
	return &CsFetchQueueLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CsFetchQueueLogic) CsFetchQueue() (resp *types.CommResp, err error) {
	// todo: add your logic here and delete this line

	return
}
