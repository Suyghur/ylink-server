package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/types/known/structpb"
	"ylink/apis/cmd/internal/svc"
	"ylink/apis/cmd/pb"
	"ylink/ext/result"
)

type PlayerFetchHistoryMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlayerFetchHistoryMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerFetchHistoryMsgLogic {
	return &PlayerFetchHistoryMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PlayerFetchHistoryMsgLogic) PlayerFetchHistoryMsg(in *pb.PlayerFetchHistoryMsgReq) (*pb.PlayerFetchHistoryMsgResp, error) {
	// todo 查询db下自己对应客服下的信息
	list, err := structpb.NewList([]interface{}{
		map[string]interface{}{
			"content":     "你好呀,我是玩家",
			"pic":         "https://www.baidu.com",
			"send_id":     "test1231",
			"receiver_id": "cs1231",
			"create_time": "2022-04-27 14:47:50",
		},
		map[string]interface{}{
			"content":     "你好呀,我是客服",
			"pic":         "",
			"send_id":     "cs1231",
			"receiver_id": "test1231",
			"create_time": "2022-04-27 14:47:50",
		},
	})
	if err != nil {
		return nil, errors.Wrap(result.NewErrMsg("fetch history message list error"), "")
	}

	return &pb.PlayerFetchHistoryMsgResp{
		TotalPage:   in.Page,
		CurrentPage: in.Page,
		List:        list,
	}, nil
}
