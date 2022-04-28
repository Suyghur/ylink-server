package cs

import (
	"context"

	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FetchHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchHistoryListLogic {
	return &FetchHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchHistoryListLogic) FetchHistoryList(req *types.CsFetchHistoryChatReq) (resp *types.CsFetchHistoryChatResp, err error) {
	// todo: add your logic here and delete this line

	return
}
