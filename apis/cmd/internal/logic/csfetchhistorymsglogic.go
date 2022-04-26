package logic

import (
	"context"

	"ylink/apis/cmd/internal/svc"
	"ylink/apis/cmd/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsFetchHistoryMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCsFetchHistoryMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsFetchHistoryMsgLogic {
	return &CsFetchHistoryMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CsFetchHistoryMsgLogic) CsFetchHistoryMsg(in *pb.CsFetchHistoryMsgReq) (*pb.CmdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CmdResp{}, nil
}
