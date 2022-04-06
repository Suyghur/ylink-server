package login

import (
	"context"

	"ylink/bff/apibff/api/internal/svc"
	"ylink/bff/apibff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCsLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsLoginLogic {
	return &CsLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CsLoginLogic) CsLogin(req *types.CsLoginInfo) (resp *types.CommResp, err error) {
	// todo: add your logic here and delete this line

	return
}
