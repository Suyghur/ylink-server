package logic

import (
	"context"
	"ylink/comm/ctxdata"
	"ylink/core/cmd/rpc/pb"

	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"

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
	gameId := ctxdata.GetGameIdFromCtx(l.ctx)
	playerId := ctxdata.GetPlayerIdFromCtx(l.ctx)
	_, err := l.svcCtx.CmdRpc.PlayerSendMsg(l.ctx, &pb.PlayerSendMsgReq{
		GameId:   gameId,
		PlayerId: playerId,
		Content:  req.Content,
		Pic:      req.Pic,
	})
	return err
}
