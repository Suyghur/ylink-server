package logic

import (
	"context"
	"ylink/comm/ctxdata"
	"ylink/core/cmd/rpc/cmd"

	"github.com/zeromicro/go-zero/core/logx"
	"ylink/bff/cmdbff/api/internal/svc"
)

type PlayerDisconnectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlayerDisconnectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerDisconnectLogic {
	return &PlayerDisconnectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlayerDisconnectLogic) PlayerDisconnect() error {
	gameId := ctxdata.GetGameIdFromCtx(l.ctx)
	playerId := ctxdata.GetPlayerIdFromCtx(l.ctx)
	_, err := l.svcCtx.CmdRpc.PlayerDisconnect(l.ctx, &cmd.PlayerDisconnectReq{
		GameId:   gameId,
		PlayerId: playerId,
	})
	return err
}
