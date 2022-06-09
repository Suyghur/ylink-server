package logic

import (
	"context"
	"ylink/core/inner/rpc/inner"

	"ylink/core/cmd/rpc/internal/svc"
	"ylink/core/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlayerDisconnectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlayerDisconnectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerDisconnectLogic {
	return &PlayerDisconnectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PlayerDisconnectLogic) PlayerDisconnect(in *pb.PlayerDisconnectReq) (*pb.PlayerDisconnectResp, error) {
	if _, err := l.svcCtx.InnerRpc.PlayerDisconnect(l.ctx, &inner.InnerPlayerDisconnectReq{GameId: in.GameId, PlayerId: in.PlayerId}); err != nil {
		return nil, err
	}
	return &pb.PlayerDisconnectResp{}, nil
}
