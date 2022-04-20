package cmd

import (
	"context"

	"ylink/bff/apibff/internal/svc"
	"ylink/bff/apibff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlayerFetchMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlayerFetchMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerFetchMsgLogic {
	return &PlayerFetchMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlayerFetchMsgLogic) PlayerFetchMsg() (resp *types.CommResp, err error) {
	// todo: add your logic here and delete this line

	return
}
