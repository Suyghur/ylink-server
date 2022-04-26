package logic

import (
	"context"
	"google.golang.org/protobuf/types/known/structpb"

	"ylink/apis/cmd/internal/svc"
	"ylink/apis/cmd/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlayerFetchCsInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlayerFetchCsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerFetchCsInfoLogic {
	return &PlayerFetchCsInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PlayerFetchCsInfoLogic) PlayerFetchCsInfo(in *pb.PlayerFetchCsInfoReq) (*pb.CmdResp, error) {
	l.Logger.Infof("invoke PlayerFetchCsInfo func, cs_id: %s", in.CsId)

	data, err := structpb.NewStruct(map[string]interface{}{
		"cs_id":         "cs_1231",
		"cs_nickname":   "vip客服1231",
		"cs_avatar_url": "https://www.baidu.com",
		"cs_signature":  "服务时间：9:30-20:30",
		"online_status": 1,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CmdResp{
		Code: 0,
		Msg:  "success",
		Data: data,
	}, nil
}
