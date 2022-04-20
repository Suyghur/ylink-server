package auth

import (
	"context"
	"ylink/apis/auth/pb"

	"ylink/bff/apibff/internal/svc"
	"ylink/bff/apibff/internal/types"

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

func (l *PlayerAuthLogic) PlayerAuth(req *types.PlayerAuthReq) (resp *types.CommResp, err error) {
	if authResp, err := l.svcCtx.AuthRpc.PlayerAuth(l.ctx, &pb.PlayerAuthReq{
		PlayerId: req.PlayerId,
		GameId:   req.GameId,
	}); err != nil {
		return &types.CommResp{
			Code: authResp.Code,
			Msg:  "success",
			Data: map[string]interface{}{},
		}, err
	} else {
		return &types.CommResp{
			Code: 0,
			Msg:  "success",
			Data: authResp.Data,
		}, nil
	}
}
