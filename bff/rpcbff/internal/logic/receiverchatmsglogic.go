package logic

import (
	"context"

	"ylink/bff/rpcbff/internal/svc"
	"ylink/bff/rpcbff/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReceiverChatMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReceiverChatMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReceiverChatMsgLogic {
	return &ReceiverChatMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReceiverChatMsgLogic) ReceiverChatMsg(in *pb.ChatMsgReq, stream pb.Rpcbff_ReceiverChatMsgServer) error {
	// todo: add your logic here and delete this line

	return nil
}
