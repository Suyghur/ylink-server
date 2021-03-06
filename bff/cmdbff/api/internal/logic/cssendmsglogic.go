package logic

import (
	"context"
	"ylink/comm/ctxdata"
	"ylink/core/cmd/rpc/cmd"

	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"

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
	csId := ctxdata.GetCsIdFromCtx(l.ctx)
	_, err := l.svcCtx.CmdRpc.CsSendMsg(l.ctx, &cmd.CsSendMsgReq{
		CsId:     csId,
		GameId:   req.GameId,
		PlayerId: req.PlayerId,
		Content:  req.Content,
		Pic:      req.Pic,
	})
	return err
}
