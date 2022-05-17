package logic

import (
	"context"
	"ylink/comm/ctxdata"
	"ylink/core/cmd/rpc/cmd"

	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsFetchHistoryMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCsFetchHistoryMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsFetchHistoryMsgLogic {
	return &CsFetchHistoryMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CsFetchHistoryMsgLogic) CsFetchHistoryMsg(req *types.CsFetchHistoryMsgReq) (resp *types.CsFetchHistoryMsgResp, err error) {
	csId := ctxdata.GetCsIdFromCtx(l.ctx)
	cmdResp, err := l.svcCtx.CmdRpc.CsFetchHistoryMsg(l.ctx, &cmd.CsFetchHistoryMsgReq{
		CsId:     csId,
		PlayerId: req.PlayerId,
		GameId:   req.GameId,
		Page:     req.Page,
		Limit:    req.Limit,
	})
	if err != nil {
		return nil, err
	}
	return &types.CsFetchHistoryMsgResp{
		TotalPage:   cmdResp.TotalPage,
		CurrentPage: cmdResp.CurrentPage,
		List:        cmdResp.List.AsSlice(),
	}, nil
}
