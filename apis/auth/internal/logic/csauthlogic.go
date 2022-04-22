package logic

import (
	"context"
	"google.golang.org/protobuf/types/known/structpb"

	"ylink/apis/auth/internal/svc"
	"ylink/apis/auth/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCsAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsAuthLogic {
	return &CsAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CsAuthLogic) CsAuth(in *pb.CsAuthReq) (*pb.AuthResp, error) {
	l.Logger.Info("invoke func CsAuth...")
	l.Logger.Infof("cs_id: %s", in.CsId)

	// todo 查询用户信息
	// todo 生成token
	if data, err := structpb.NewStruct(map[string]interface{}{
		"token": "cs_auth",
	}); err != nil {
		return &pb.AuthResp{
			Code: 1,
			Msg:  err.Error(),
			Data: nil,
		}, err
	} else {
		return &pb.AuthResp{
			Code: 0,
			Msg:  "success",
			Data: data,
		}, nil
	}
}
