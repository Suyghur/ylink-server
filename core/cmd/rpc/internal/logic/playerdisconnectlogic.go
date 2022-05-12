package logic

import (
	"context"
	"ylink/core/cmd/rpc/internal/svc"
	"ylink/core/cmd/rpc/pb"
	"ylink/core/inner/rpc/inner"

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
	// 调用inner服务玩家状态
	_, err := l.svcCtx.InnerRpc.PlayerDisconnect(l.ctx, &inner.InnerPlayerDisconnectReq{
		PlayerId: in.PlayerId,
		GameId:   in.GameId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.PlayerDisconnectResp{}, nil
}
