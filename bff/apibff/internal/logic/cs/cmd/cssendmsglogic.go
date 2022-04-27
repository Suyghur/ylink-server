package cmd

import (
	"context"

	"ylink/bff/apibff/internal/svc"
	"ylink/bff/apibff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsSendMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCsSendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsSendMsgLogic {
	return &CsSendMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CsSendMsgLogic) CsSendMsg(req *types.CsSendMsgReq) error {
	// todo: add your logic here and delete this line

	return nil
}
