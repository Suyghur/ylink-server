package logic

import (
	"context"
	"ylink/core/cmd/rpc/internal/svc"
	"ylink/core/cmd/rpc/pb"
	"ylink/core/inner/rpc/inner"

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
	innerResp, err := l.svcCtx.InnerRpc.PlayerFetchCsInfo(l.ctx, &inner.InnerPlayerFetchCsInfoReq{
		PlayerId: in.PlayerId,
		GameId:   in.GameId,
		CsId:     in.CsId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.PlayerFetchCsInfoResp{
		CsId:         innerResp.CsId,
		CsNickname:   innerResp.CsNickname,
		CsAvatarUrl:  innerResp.CsAvatarUrl,
		CsSignature:  innerResp.CsSignature,
		OnlineStatus: innerResp.OnlineStatus,
	}, nil
}
