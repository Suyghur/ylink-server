package logic

import (
	"context"

	"ylink/apis/cmd/internal/svc"
	"ylink/apis/cmd/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlayerFetchHistoryMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlayerFetchHistoryMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerFetchHistoryMsgLogic {
	return &PlayerFetchHistoryMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PlayerFetchHistoryMsgLogic) PlayerFetchHistoryMsg(in *pb.PlayerFetchHistoryMsgReq) (*pb.CmdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CmdResp{}, nil
}
