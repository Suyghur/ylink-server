package logic

import (
	"context"
	"ylink/core/cmd/rpc/internal/svc"
	"ylink/core/cmd/rpc/pb"
	"ylink/core/inner/rpc/inner"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsFetchPlayerQueueLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCsFetchPlayerQueueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsFetchPlayerQueueLogic {
	return &CsFetchPlayerQueueLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CsFetchPlayerQueueLogic) CsFetchPlayerQueue(in *pb.CsFetchPlayerQueueReq) (*pb.CsFetchPlayerQueueResp, error) {
	// 调用inner服务查询等待用户的队列
	innerResp, err := l.svcCtx.InnerRpc.CsFetchPlayerQueue(l.ctx, &inner.InnerCsFetchPlayerQueueReq{
		Limit: in.Limit,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CsFetchPlayerQueueResp{
		List: innerResp.List,
	}, nil
}
