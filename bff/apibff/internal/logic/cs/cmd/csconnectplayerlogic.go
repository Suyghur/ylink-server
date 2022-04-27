package cmd

import (
	"context"
	"ylink/apis/cmd/pb"

	"ylink/bff/apibff/internal/svc"
	"ylink/bff/apibff/internal/types"

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
	_, err := l.svcCtx.CmdRpc.CsConnectPlayer(l.ctx, &pb.CsConnectPlayerReq{
		PlayerId: req.PlayerId,
		GameId:   req.GameId,
		CsId:     "",
	})
	if err != nil {
		return err
	}
	return nil
}
