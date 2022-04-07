package logic

import (
	"context"

	"ylink/gateway/rpc/internal/svc"
	"ylink/gateway/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCsLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsLoginLogic {
	return &CsLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CsLoginLogic) CsLogin(in *pb.CsLoginReq) (*pb.LoginResp, error) {
	// todo: add your logic here and delete this line

	return &pb.LoginResp{}, nil
}
