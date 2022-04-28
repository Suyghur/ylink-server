package player

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"ylink/bff/cmdbff/api/internal/svc"
)

type DisconnectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDisconnectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DisconnectLogic {
	return &DisconnectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DisconnectLogic) Disconnect() error {
	// todo: add your logic here and delete this line

	return nil
}
