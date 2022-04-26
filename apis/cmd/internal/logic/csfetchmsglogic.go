package logic

import (
	"context"

	"ylink/apis/cmd/internal/svc"
	"ylink/apis/cmd/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsFetchMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCsFetchMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsFetchMsgLogic {
	return &CsFetchMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CsFetchMsgLogic) CsFetchMsg(in *pb.CsFetchMsgReq) (*pb.CmdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CmdResp{}, nil
}
