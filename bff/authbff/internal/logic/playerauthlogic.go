package logic

import (
	"context"
	"ylink/apis/auth/pb"

	"ylink/bff/authbff/internal/svc"
	"ylink/bff/authbff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlayerAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlayerAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerAuthLogic {
	return &PlayerAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlayerAuthLogic) PlayerAuth(req *types.PlayerAuthReq) (resp *types.AuthResp, err error) {
	if authResp, err := l.svcCtx.AuthRpc.PlayerAuth(l.ctx, &pb.PlayerAuthReq{
		PlayerId: req.PlayerId,
		GameId:   req.GameId,
	}); err != nil {
		return &types.AuthResp{
			Code: authResp.Code,
			Msg:  authResp.Msg,
			Data: map[string]interface{}{},
		}, err
	} else {
		return &types.AuthResp{
			Code: authResp.Code,
			Msg:  authResp.Msg,
			Data: authResp.Data,
		}, nil
	}
}
