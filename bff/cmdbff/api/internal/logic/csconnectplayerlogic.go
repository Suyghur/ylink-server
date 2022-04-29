package logic

import (
	"context"
	"ylink/core/cmd/rpc/cmd"
	"ylink/ext/ctxdata"

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
		PlayerId: req.PlayerId,
		GameId:   req.GameId,
	})
	if err != nil {
		return err
	}
	return nil
}
