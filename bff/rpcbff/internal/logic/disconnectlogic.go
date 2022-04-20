package logic

import (
	"context"
	"google.golang.org/protobuf/types/known/structpb"

	"ylink/bff/rpcbff/internal/svc"
	"ylink/bff/rpcbff/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DisconnectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDisconnectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DisconnectLogic {
	return &DisconnectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DisconnectLogic) Disconnect(in *pb.CommandReq) (*pb.CommandResp, error) {
	// todo: 验证token
	// todo: 把关联的stream从资源pool中移除

	l.Logger.Info("invoke func disconnect")
	l.Logger.Infof("%s", in.Token)

	data, _ := structpb.NewStruct(map[string]interface{}{})
	return &pb.CommandResp{
		CommandCode: 0,
		CommandMsg:  "success",
		CommandData: data,
	}, nil
}
