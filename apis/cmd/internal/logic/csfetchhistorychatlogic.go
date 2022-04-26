package logic

import (
	"context"

	"ylink/apis/cmd/internal/svc"
	"ylink/apis/cmd/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsFetchHistoryChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCsFetchHistoryChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsFetchHistoryChatLogic {
	return &CsFetchHistoryChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CsFetchHistoryChatLogic) CsFetchHistoryChat(in *pb.CsFetchHistoryChatReq) (*pb.CmdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CmdResp{}, nil
}
