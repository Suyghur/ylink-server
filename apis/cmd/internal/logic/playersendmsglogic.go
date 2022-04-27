package logic

import (
	"context"

	"ylink/apis/cmd/internal/svc"
	"ylink/apis/cmd/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlayerSendMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlayerSendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerSendMsgLogic {
	return &PlayerSendMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PlayerSendMsgLogic) PlayerSendMsg(in *pb.PlayerSendMsgReq) (*pb.PlayerSendMsgResp, error) {
	// todo 投递到对应客服的收件箱
	// todo 写入db

	return &pb.PlayerSendMsgResp{}, nil
}
