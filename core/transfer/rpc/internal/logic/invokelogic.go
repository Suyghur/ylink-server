package logic

import (
	"context"

	"ylink/core/transfer/rpc/internal/svc"
	"ylink/core/transfer/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type InvokeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInvokeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InvokeLogic {
	return &InvokeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InvokeLogic) Invoke(in *pb.TransferReq) (*pb.TransferResp, error) {
	// todo: add your logic here and delete this line

	return &pb.TransferResp{}, nil
}
