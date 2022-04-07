package logic

import (
	"context"

	"ylink/gateway/rpc/internal/svc"
	"ylink/gateway/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlayerLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlayerLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerLoginLogic {
	return &PlayerLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PlayerLoginLogic) PlayerLogin(in *pb.PlayerLoginReq) (*pb.LoginResp, error) {
	// todo: add your logic here and delete this line

	return &pb.LoginResp{
		AccessToken:  in.PlayerId,
		AccessExpire: 100,
		RefreshAfter: 100,
		Url:          "www.baidu.com"}, nil
}
