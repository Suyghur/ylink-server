package cs

import (
	"context"

	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FetchMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchMsgLogic {
	return &FetchMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchMsgLogic) FetchMsg(req *types.CsFetchMsgReq) (resp *types.CsFetchMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
