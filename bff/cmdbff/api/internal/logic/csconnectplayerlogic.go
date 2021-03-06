package logic

import (
	"context"
	"ylink/comm/ctxdata"
	"ylink/core/cmd/rpc/cmd"

	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsConnectPlayerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCsConnectPlayerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsConnectPlayerLogic {
	return &CsConnectPlayerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CsConnectPlayerLogic) CsConnectPlayer(req *types.CsConnectPlayerReq) error {
	csId := ctxdata.GetCsIdFromCtx(l.ctx)
	_, err := l.svcCtx.CmdRpc.CsConnectPlayer(l.ctx, &cmd.CsConnectPlayerReq{
		CsId:     csId,
		GameId:   req.GameId,
		PlayerId: req.PlayerId,
	})
	if err != nil {
		return err
	}
	return nil
}
