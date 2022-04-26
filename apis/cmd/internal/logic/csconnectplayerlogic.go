package logic

import (
	"context"

	"ylink/apis/cmd/internal/svc"
	"ylink/apis/cmd/pb"

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

func (l *CsConnectPlayerLogic) CsConnectPlayer(in *pb.CsConnectPlayerReq) (*pb.CmdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CmdResp{}, nil
}
