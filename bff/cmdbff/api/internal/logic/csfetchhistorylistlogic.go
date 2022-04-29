package logic

import (
	"context"
	"ylink/core/cmd/rpc/cmd"
	"ylink/ext/ctxdata"

	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsFetchHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCsFetchHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsFetchHistoryListLogic {
	return &CsFetchHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CsFetchHistoryListLogic) CsFetchHistoryList(req *types.CsFetchHistoryChatReq) (resp *types.CsFetchHistoryChatResp, err error) {
	csId := ctxdata.GetCsIdFromCtx(l.ctx)
	cmdResp, err := l.svcCtx.CmdRpc.CsFetchHistoryChat(l.ctx, &cmd.CsFetchHistoryChatReq{
		CsId:  csId,
		Page:  req.Page,
		Limit: req.Limit,
	})
	if err != nil {
		return nil, err
	}
	return &types.CsFetchHistoryChatResp{
		TotalPage:   cmdResp.TotalPage,
		CurrentPage: cmdResp.CurrentPage,
		List:        cmdResp.List.AsSlice(),
	}, nil
}
