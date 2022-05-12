package logic

import (
	"context"
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
	// todo: 建立玩家-客服的映射关系

	return &pb.InnerCsConnectPlayerResp{}, nil
}
