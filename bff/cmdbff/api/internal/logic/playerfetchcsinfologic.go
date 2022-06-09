package logic

import (
	"context"
	"ylink/comm/ctxdata"
	"ylink/core/cmd/rpc/cmd"

	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlayerFetchCsInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlayerFetchCsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerFetchCsInfoLogic {
	return &PlayerFetchCsInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlayerFetchCsInfoLogic) PlayerFetchCsInfo(req *types.PlayerFetchCsInfoReq) (resp *types.PlayerFetchCsInfoResp, err error) {
	gameId := ctxdata.GetGameIdFromCtx(l.ctx)
	playerId := ctxdata.GetPlayerIdFromCtx(l.ctx)
	cmdResp, err := l.svcCtx.CmdRpc.PlayerFetchCsInfo(l.ctx, &cmd.PlayerFetchCsInfoReq{
		GameId:   gameId,
		PlayerId: playerId,
		CsId:     req.CsId,
	})
	if err != nil {
		return nil, err
	}
	return &types.PlayerFetchCsInfoResp{
		CsId:         cmdResp.CsId,
		CsNickname:   cmdResp.CsNickname,
		CsAvatarUrl:  cmdResp.CsAvatarUrl,
		CsSignature:  cmdResp.CsSignature,
		OnlineStatus: cmdResp.OnlineStatus,
	}, nil
}
