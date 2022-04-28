package cs

import (
	"context"
	"ylink/core/cmd/rpc/cmd"
	"ylink/ext/ctxdata"

	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FetchHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchHistoryListLogic {
	return &FetchHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchHistoryListLogic) FetchHistoryList(req *types.CsFetchHistoryChatReq) (resp *types.CsFetchHistoryChatResp, err error) {
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
