package cs

import (
	"context"

	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsFetchHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCsFetchHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsFetchHistoryListLogic {
	return &CsFetchHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CsFetchHistoryListLogic) CsFetchHistoryList(req *types.CsFetchHistoryChatReq) (resp *types.CsFetchHistoryChatResp, err error) {
	// todo: add your logic here and delete this line

	return
}
