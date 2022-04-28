package cs

import (
	"context"

	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FetchPlayerQueueLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchPlayerQueueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchPlayerQueueLogic {
	return &FetchPlayerQueueLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchPlayerQueueLogic) FetchPlayerQueue(req *types.CsFetchPlayerQueueReq) (resp *types.CsFetchPlayerQueueResp, err error) {
	// todo: add your logic here and delete this line

	return
}
