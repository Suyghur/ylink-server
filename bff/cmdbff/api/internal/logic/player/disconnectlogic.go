package player

import (
	"context"
	"ylink/core/cmd/rpc/cmd"
	"ylink/ext/ctxdata"

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
	playerId := ctxdata.GetPlayerIdFromCtx(l.ctx)
	gameId := ctxdata.GetGameIdFromCtx(l.ctx)
	_, err := l.svcCtx.CmdRpc.PlayerDisconnect(l.ctx, &cmd.PlayerDisconnectReq{
		PlayerId: playerId,
		GameId:   gameId,
	})
	return err
}
