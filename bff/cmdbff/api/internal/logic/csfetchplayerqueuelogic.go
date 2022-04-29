package logic

import (
	"context"
	"ylink/core/cmd/rpc/cmd"

	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsFetchPlayerQueueLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCsFetchPlayerQueueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsFetchPlayerQueueLogic {
	return &CsFetchPlayerQueueLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CsFetchPlayerQueueLogic) CsFetchPlayerQueue(req *types.CsFetchPlayerQueueReq) (resp *types.CsFetchPlayerQueueResp, err error) {
	cmdResp, err := l.svcCtx.CmdRpc.CsFetchPlayerQueue(l.ctx, &cmd.CsFetchPlayerQueueReq{
		Limit: req.Limit,
	})
	if err != nil {
		return nil, err
	}
	return &types.CsFetchPlayerQueueResp{
		List: cmdResp.List.AsSlice(),
	}, nil
}
