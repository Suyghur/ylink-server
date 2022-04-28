package logic

import (
	"context"

	"ylink/core/cmd/rpc/internal/svc"
	"ylink/core/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlayerFetchCsInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlayerFetchCsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerFetchCsInfoLogic {
	return &PlayerFetchCsInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PlayerFetchCsInfoLogic) PlayerFetchCsInfo(in *pb.PlayerFetchCsInfoReq) (*pb.PlayerFetchCsInfoResp, error) {
	return &pb.PlayerFetchCsInfoResp{
		CsId:         in.CsId,
		CsNickname:   "vip客服1231",
		CsAvatarUrl:  "https://www.baiduc.om",
		CsSignature:  "服务时间：9:30-20:30",
		OnlineStatus: 1,
	}, nil
}
