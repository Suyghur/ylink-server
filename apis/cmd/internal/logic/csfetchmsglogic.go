package logic

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/structpb"
	"ylink/ext/result"

	"ylink/apis/cmd/internal/svc"
	"ylink/apis/cmd/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsFetchMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCsFetchMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsFetchMsgLogic {
	return &CsFetchMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CsFetchMsgLogic) CsFetchMsg(in *pb.CsFetchMsgReq) (*pb.CsFetchMsgResp, error) {
	// todo 查询自己收件箱下对应玩家的信息
	list, err := structpb.NewList([]interface{}{
		map[string]interface{}{
			"content":     "你好呀,我是玩家",
			"pic":         "https://www.baidu.com",
			"send_id":     in.PlayerId,
			"receiver_id": in.CsId,
			"create_time": "2022-04-27 14:47:50",
		},
		map[string]interface{}{
			"content":     "有个问题需要帮忙处理一下",
			"pic":         "",
			"send_id":     in.PlayerId,
			"receiver_id": in.CsId,
			"create_time": "2022-04-27 14:47:50",
		},
	})
	if err != nil {
		return nil, errors.Wrap(result.NewErrMsg("fetch cs message list error"), "")
	}
	return &pb.CsFetchMsgResp{
		List: list,
	}, nil
}
