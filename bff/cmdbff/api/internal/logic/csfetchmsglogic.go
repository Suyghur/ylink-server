package logic

import (
	"context"
	"ylink/core/cmd/rpc/cmd"
	"ylink/ext/ctxdata"

	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsFetchMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCsFetchMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsFetchMsgLogic {
	return &CsFetchMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CsFetchMsgLogic) CsFetchMsg(req *types.CsFetchMsgReq) (resp *types.CsFetchMsgResp, err error) {
	csId := ctxdata.GetCsIdFromCtx(l.ctx)
	cmdResp, err := l.svcCtx.CmdRpc.CsFetchMsg(l.ctx, &cmd.CsFetchMsgReq{
		CsId:     csId,
		PlayerId: req.PlayerId,
		GameId:   req.GameId,
	})
	if err != nil {
		return nil, err
	}
	return &types.CsFetchMsgResp{
		List: cmdResp.List.AsSlice(),
	}, nil
}
