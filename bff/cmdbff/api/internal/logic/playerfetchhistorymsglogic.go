package logic

import (
	"context"
	"ylink/comm/ctxdata"
	"ylink/core/cmd/rpc/pb"

	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlayerFetchHistoryMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlayerFetchHistoryMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerFetchHistoryMsgLogic {
	return &PlayerFetchHistoryMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlayerFetchHistoryMsgLogic) PlayerFetchHistoryMsg(req *types.PlayerFetchHistoryMsgReq) (resp *types.PlayerFetchHistoryMsgResp, err error) {
	gameId := ctxdata.GetGameIdFromCtx(l.ctx)
	playerId := ctxdata.GetPlayerIdFromCtx(l.ctx)
	cmdResp, err := l.svcCtx.CmdRpc.PlayerFetchHistoryMsg(l.ctx, &pb.PlayerFetchHistoryMsgReq{
		GameId:   gameId,
		PlayerId: playerId,
		Page:     req.Page,
		Limit:    req.Limit,
	})
	if err != nil {
		return nil, err
	}
	return &types.PlayerFetchHistoryMsgResp{
		TotalPage:   cmdResp.TotalPage,
		CurrentPage: cmdResp.CurrentPage,
		List:        cmdResp.List.AsSlice(),
	}, nil
}
