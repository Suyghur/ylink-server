package cmd

import (
	"context"

	"ylink/bff/apibff/api/internal/svc"
	"ylink/bff/apibff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsFetchMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCsFetchMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsFetchMsgLogic {
	return &CsFetchMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CsFetchMsgLogic) CsFetchMsg(req *types.CsFetchMsgReq) (resp *types.CommResp, err error) {
	// todo: add your logic here and delete this line

	return
}
