package cmd

import (
	"context"
	"ylink/apis/cmd/pb"

	"ylink/bff/apibff/internal/svc"
	"ylink/bff/apibff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlayerFetchCsInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlayerFetchCsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerFetchCsInfoLogic {
	return &PlayerFetchCsInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlayerFetchCsInfoLogic) PlayerFetchCsInfo(req *types.PlayerFetchCsInfoReq) (resp *types.PlayerFetchCsInfoResp, err error) {
	cmdResp, err := l.svcCtx.CmdRpc.PlayerFetchCsInfo(l.ctx, &pb.PlayerFetchCsInfoReq{
		CsId: req.CsId,
	})
	if err != nil {
		return nil, err
	}
	return &types.PlayerFetchCsInfoResp{
		CsId:         cmdResp.CsId,
		CsNickname:   cmdResp.CsNickname,
		CsAvatarUrl:  cmdResp.CsAvatarUrl,
		CsSignature:  cmdResp.CsSignature,
		OnlineStatus: cmdResp.OnlineStatus,
	}, nil
}
