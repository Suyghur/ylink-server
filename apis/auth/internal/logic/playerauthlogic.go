package logic

import (
	"context"
	"google.golang.org/protobuf/types/known/structpb"

	"ylink/apis/auth/internal/svc"
	"ylink/apis/auth/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlayerAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlayerAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerAuthLogic {
	return &PlayerAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PlayerAuthLogic) PlayerAuth(in *pb.PlayerAuthReq) (*pb.AuthResp, error) {
	l.Logger.Info("invoke func PlayerAuth...")
	l.Logger.Infof("player_id: %s", in.PlayerId)
	l.Logger.Infof("game_id: %s", in.GameId)

	// todo 查询用户信息
	// todo 生成token
	if data, err := structpb.NewStruct(map[string]interface{}{
		"token": "player_auth",
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
