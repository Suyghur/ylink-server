package player

import (
	"context"
	"ylink/core/cmd/rpc/cmd"
	"ylink/ext/ctxdata"

	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FetchCsInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchCsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchCsInfoLogic {
	return &FetchCsInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchCsInfoLogic) FetchCsInfo(req *types.PlayerFetchCsInfoReq) (resp *types.PlayerFetchCsInfoResp, err error) {
	playerId := ctxdata.GetPlayerIdFromCtx(l.ctx)
	gameId := ctxdata.GetGameIdFromCtx(l.ctx)
	l.Logger.Infof("player id: %s", playerId)
	cmdResp, err := l.svcCtx.CmdRpc.PlayerFetchCsInfo(l.ctx, &cmd.PlayerFetchCsInfoReq{
		PlayerId: playerId,
		GameId:   gameId,
		CsId:     req.CsId,
	})
	if err != nil {
		l.Logger.Info(err.Error())
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
