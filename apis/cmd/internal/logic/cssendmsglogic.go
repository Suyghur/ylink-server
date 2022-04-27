package logic

import (
	"context"
	"ylink/apis/cmd/internal/svc"
	"ylink/apis/cmd/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsSendMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCsSendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsSendMsgLogic {
	return &CsSendMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CsSendMsgLogic) CsSendMsg(in *pb.CsSendMsgReq) (*pb.CsSendMsgResp, error) {
	// todo 投递到对应客服的收件箱
	// todo 写入db
	return &pb.CsSendMsgResp{}, nil
}
