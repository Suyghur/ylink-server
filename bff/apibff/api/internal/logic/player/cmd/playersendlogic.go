package cmd

import (
	"context"

	"ylink/bff/apibff/api/internal/svc"
	"ylink/bff/apibff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlayerSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlayerSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerSendLogic {
	return &PlayerSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlayerSendLogic) PlayerSend(req *types.SendReq) (resp *types.CommResp, err error) {
	// todo: add your logic here and delete this line

	return
}
