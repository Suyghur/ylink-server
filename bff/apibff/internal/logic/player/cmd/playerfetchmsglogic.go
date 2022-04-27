package cmd

import (
	"context"
	"ylink/apis/cmd/pb"
	"ylink/ext/ctxdata"

	"ylink/bff/apibff/internal/svc"
	"ylink/bff/apibff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlayerFetchMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlayerFetchMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerFetchMsgLogic {
	return &PlayerFetchMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlayerFetchMsgLogic) PlayerFetchMsg() (resp *types.PlayerFetchMsgResp, err error) {
	playerId := ctxdata.GetGameIdFromCtx(l.ctx)
	gameId := ctxdata.GetGameIdFromCtx(l.ctx)
	cmdResp, err := l.svcCtx.CmdRpc.PlayerFetchMsg(l.ctx, &pb.PlayerFetchMsgReq{
		PlayerId: playerId,
		GameId:   gameId,
	})
	if err != nil {
		return nil, err
	}
	return &types.PlayerFetchMsgResp{
		List: cmdResp.List.AsSlice(),
	}, nil
}
