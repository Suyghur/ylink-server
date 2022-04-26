package logic

import (
	"context"

	"ylink/apis/cmd/internal/svc"
	"ylink/apis/cmd/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsFetchPlayerQueueLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCsFetchPlayerQueueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsFetchPlayerQueueLogic {
	return &CsFetchPlayerQueueLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CsFetchPlayerQueueLogic) CsFetchPlayerQueue(in *pb.CsFetchPlayerQueueReq) (*pb.CmdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CmdResp{}, nil
}
