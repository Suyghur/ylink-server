package cmd

import (
	"context"

	"ylink/bff/apibff/internal/svc"
	"ylink/bff/apibff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlayerFetchHistoryMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlayerFetchHistoryMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerFetchHistoryMsgLogic {
	return &PlayerFetchHistoryMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlayerFetchHistoryMsgLogic) PlayerFetchHistoryMsg(req *types.PlayerFetchHistoryMsgReq) (resp *types.CommResp, err error) {
	// todo: add your logic here and delete this line

	return
}
