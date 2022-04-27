package cmd

import (
	"context"
	"ylink/apis/cmd/pb"
	"ylink/ext/ctxdata"

	"ylink/bff/apibff/internal/svc"
	"ylink/bff/apibff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlayerSendMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlayerSendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerSendMsgLogic {
	return &PlayerSendMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlayerSendMsgLogic) PlayerSendMsg(req *types.PlayerSendMsgReq) error {
	playerId := ctxdata.GetPlayerIdFromCtx(l.ctx)
	gameId := ctxdata.GetGameIdFromCtx(l.ctx)
	_, err := l.svcCtx.CmdRpc.PlayerSendMsg(l.ctx, &pb.PlayerSendMsgReq{
		PlayerId: playerId,
		GameId:   gameId,
		Content:  req.Content,
		Pic:      req.Pic,
	})
	return err
}
