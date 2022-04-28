package logic

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/structpb"
	"ylink/ext/result"

	"ylink/core/cmd/rpc/internal/svc"
	"ylink/core/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlayerFetchMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlayerFetchMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerFetchMsgLogic {
	return &PlayerFetchMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PlayerFetchMsgLogic) PlayerFetchMsg(in *pb.PlayerFetchMsgReq) (*pb.PlayerFetchMsgResp, error) {
	// todo 全量取出自己收件箱下的信息
	list, err := structpb.NewList([]interface{}{
		map[string]interface{}{
			"content":     "你好呀,我是玩家",
			"pic":         "https://www.baidu.com",
			"send_id":     in.PlayerId,
			"receiver_id": "cs1231",
			"create_time": "2022-04-27 14:47:50",
		},
		map[string]interface{}{
			"content":     "你好呀,我是客服",
			"pic":         "",
			"send_id":     "cs1231",
			"receiver_id": in.PlayerId,
			"create_time": "2022-04-27 14:47:50",
		},
	})
	if err != nil {
		return nil, errors.Wrap(result.NewErrMsg("fetch message list error"), "")
	}

	return &pb.PlayerFetchMsgResp{
		List: list,
	}, nil
}
