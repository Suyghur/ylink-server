package player

import (
	"context"

	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FetchHistoryMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchHistoryMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchHistoryMsgLogic {
	return &FetchHistoryMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchHistoryMsgLogic) FetchHistoryMsg(req *types.PlayerFetchHistoryMsgReq) (resp *types.PlayerFetchHistoryMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
