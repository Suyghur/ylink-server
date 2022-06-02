package logic

import (
	"context"
	treemap "github.com/liyue201/gostl/ds/map"
	"ylink/core/inner/rpc/internal/ext"
	"ylink/core/inner/rpc/internal/svc"
	"ylink/core/inner/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsConnectPlayerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCsConnectPlayerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsConnectPlayerLogic {
	return &CsConnectPlayerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CsConnectPlayerLogic) CsConnectPlayer(in *pb.InnerCsConnectPlayerReq) (*pb.InnerCsConnectPlayerResp, error) {
	if ext.GameConnectedMap.Contains(in.GameId) {
		playerConnMap := ext.GameConnectedMap.Get(in.GameId).(*treemap.Map)
		playerConnMap.Insert(in.PlayerId, in.CsId)
	} else {
		playerConnMap := treemap.New(treemap.WithGoroutineSafe())
		playerConnMap.Insert(in.PlayerId, in.CsId)
		ext.GameConnectedMap.Insert(in.GameId, playerConnMap)
	}

	return &pb.InnerCsConnectPlayerResp{}, nil
}
