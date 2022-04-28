package cs

import (
	"context"

	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConnectPlayerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConnectPlayerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConnectPlayerLogic {
	return &ConnectPlayerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConnectPlayerLogic) ConnectPlayer(req *types.CsConnectPlayerReq) error {
	// todo: add your logic here and delete this line

	return nil
}
