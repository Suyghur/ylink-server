package logic

import (
	"context"
	"ylink/core/cmd/rpc/internal/svc"
	"ylink/core/cmd/rpc/pb"
	"ylink/core/inner/rpc/inner"

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

func (l *CsConnectPlayerLogic) CsConnectPlayer(in *pb.CsConnectPlayerReq) (*pb.CsConnectPlayerResp, error) {
	// 调用inner服务建立映射关系
	_, err := l.svcCtx.InnerRpc.CsConnectPlayer(l.ctx, &inner.InnerCsConnectPlayerReq{
		CsId:     in.CsId,
		GameId:   in.GameId,
		PlayerId: in.PlayerId,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CsConnectPlayerResp{}, nil
}
