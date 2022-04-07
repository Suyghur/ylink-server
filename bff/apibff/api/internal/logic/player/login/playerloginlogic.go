package login

import (
	"context"
	"ylink/gateway/rpc/gateway"

	"ylink/bff/apibff/api/internal/svc"
	"ylink/bff/apibff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *PlayerLoginLogic) PlayerLogin(req *types.PlayerLoginInfo) (resp *types.CommResp, err error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.GatewayRpc.PlayerLogin(l.ctx, &gateway.PlayerLoginReq{
		PlayerId: req.PlayerId,
		GameId:   req.GameId,
	})
	if err != nil {
		return &types.CommResp{Code: -1, Msg: err.Error(), Data: map[string]interface{}{}}, err
	}
	return &types.CommResp{Code: 0, Msg: "success", Data: map[string]interface{}{
		"access_token":  rpcResp.AccessToken,
		"access_expire": rpcResp.AccessExpire,
		"refreshAfter":  rpcResp.RefreshAfter,
		"url":           rpcResp.Url,
	}}, nil
}
