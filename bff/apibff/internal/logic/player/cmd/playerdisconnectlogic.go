package cmd

import (
	"context"
	"ylink/apis/cmd/pb"
	"ylink/ext/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
	"ylink/bff/apibff/internal/svc"
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
	playerId := ctxdata.GetPlayerIdFromCtx(l.ctx)
	gameId := ctxdata.GetGameIdFromCtx(l.ctx)
	_, err := l.svcCtx.CmdRpc.PlayerDisconnect(l.ctx, &pb.PlayerDisconnectReq{
		PlayerId: playerId,
		GameId:   gameId,
	})
	return err
}
