package logic

import (
	"context"
	"ylink/core/cmd/rpc/pb"
	"ylink/ext/ctxdata"

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
	playerId := ctxdata.GetPlayerIdFromCtx(l.ctx)
	gameId := ctxdata.GetGameIdFromCtx(l.ctx)
	cmdResp, err := l.svcCtx.CmdRpc.PlayerFetchHistoryMsg(l.ctx, &pb.PlayerFetchHistoryMsgReq{
		PlayerId: playerId,
		GameId:   gameId,
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
