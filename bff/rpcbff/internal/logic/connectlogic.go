package logic

import (
	"context"
	"google.golang.org/protobuf/types/known/structpb"
	"ylink/apis/auth/auth"

	"github.com/zeromicro/go-zero/core/logx"
	"ylink/bff/rpcbff/internal/svc"
	"ylink/bff/rpcbff/pb"
)

type ConnectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConnectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConnectLogic {
	return &ConnectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConnectLogic) Connect(in *pb.CommandReq, stream pb.Rpcbff_ConnectServer) error {
	// todo: 验证token
	// todo: 把stream放入资源pool
	if authResp, err := l.svcCtx.AuthRpc.CheckAuth(l.ctx, &auth.CheckAuthReq{
		AccessToken: in.AccessToken,
	}); err != nil {
		return stream.Send(&pb.CommandResp{
			CommandCode: authResp.Code,
			CommandMsg:  authResp.Msg,
			CommandData: authResp.Data,
		})
	}

	data, _ := structpb.NewStruct(map[string]interface{}{})

	return stream.Send(&pb.CommandResp{
		CommandCode: 0,
		CommandMsg:  "success",
		CommandData: data,
	})
}
