package login

import (
	"context"

	"ylink/bff/apibff/api/internal/svc"
	"ylink/bff/apibff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlayerLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlayerLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerLoginLogic {
	return &PlayerLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlayerLoginLogic) PlayerLogin(req *types.PlayerInfo) (resp *types.CommResp, err error) {
	// todo: add your logic here and delete this line

	return
}
