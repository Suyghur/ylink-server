package logic

import (
	"context"
	"ylink/core/inner/rpc/internal/ext"

	"ylink/core/inner/rpc/internal/svc"
	"ylink/core/inner/rpc/pb"

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

func (l *PlayerDisconnectLogic) PlayerDisconnect(in *pb.InnerPlayerDisconnectReq) (*pb.InnerPlayerDisconnectResp, error) {
	ext.RemoveConnectedPlayerInfo(in.GameId, in.PlayerId)
	return &pb.InnerPlayerDisconnectResp{}, nil
}
