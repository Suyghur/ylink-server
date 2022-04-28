package cs

import (
	"context"
	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"
	"ylink/core/cmd/rpc/cmd"
	"ylink/ext/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConnectPlayerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConnectPlayerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConnectPlayerLogic {
	return &ConnectPlayerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConnectPlayerLogic) ConnectPlayer(req *types.CsConnectPlayerReq) error {
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
