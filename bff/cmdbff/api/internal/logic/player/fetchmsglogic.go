package player

import (
	"context"
	"ylink/core/cmd/rpc/pb"
	"ylink/ext/ctxdata"

	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FetchMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchMsgLogic {
	return &FetchMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchMsgLogic) FetchMsg() (resp *types.PlayerFetchMsgResp, err error) {
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
