package cs

import (
	"context"

	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsFetchHistoryMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCsFetchHistoryMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsFetchHistoryMsgLogic {
	return &CsFetchHistoryMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CsFetchHistoryMsgLogic) CsFetchHistoryMsg(req *types.CsFetchHistoryMsgReq) (resp *types.CsFetchHistoryMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
