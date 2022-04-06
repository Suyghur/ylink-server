package cmd

import (
	"context"

	"ylink/bff/apibff/api/internal/svc"
	"ylink/bff/apibff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlayerLogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlayerLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerLogoutLogic {
	return &PlayerLogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlayerLogoutLogic) PlayerLogout() (resp *types.CommResp, err error) {
	// todo: add your logic here and delete this line

	return
}
