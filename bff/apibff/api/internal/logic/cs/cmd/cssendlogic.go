package cmd

import (
	"context"

	"ylink/bff/apibff/api/internal/svc"
	"ylink/bff/apibff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCsSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsSendLogic {
	return &CsSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CsSendLogic) CsSend(req *types.ChatMsgReq) (resp *types.CommResp, err error) {
	// todo: add your logic here and delete this line

	return
}
