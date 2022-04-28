package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"ylink/bff/authbff/api/internal/svc"
	"ylink/bff/authbff/api/internal/types"
	"ylink/core/auth/rpc/auth"
)

type PlayerLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlayerLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerLoginLogic {
	return &PlayerLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlayerLoginLogic) PlayerLogin(req *types.PlayerAuthReq) (resp *types.AuthResp, err error) {
	authResp, err := l.svcCtx.AuthRpc.PlayerAuth(l.ctx, &auth.PlayerAuthReq{
		PlayerId: req.PlayerId,
		GameId:   req.GameId,
	})
	if err != nil {
		return nil, err
	}
	return &types.AuthResp{
		AccessToken: authResp.AccessToken,
	}, nil
}
