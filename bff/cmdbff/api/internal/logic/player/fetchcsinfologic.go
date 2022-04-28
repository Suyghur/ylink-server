package player

import (
	"context"

	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FetchCsInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchCsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchCsInfoLogic {
	return &FetchCsInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchCsInfoLogic) FetchCsInfo(req *types.PlayerFetchCsInfoReq) (resp *types.PlayerFetchCsInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
