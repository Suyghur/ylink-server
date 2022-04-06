package cmd

import (
	"context"

	"ylink/bff/apibff/api/internal/svc"
	"ylink/bff/apibff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlayerConnectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlayerConnectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerConnectLogic {
	return &PlayerConnectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlayerConnectLogic) PlayerConnect(req *types.PlayerConnectReq) (resp *types.CommResp, err error) {
	// todo: add your logic here and delete this line

	return
}
