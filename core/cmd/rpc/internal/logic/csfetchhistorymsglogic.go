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

type CsFetchHistoryMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCsFetchHistoryMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsFetchHistoryMsgLogic {
	return &CsFetchHistoryMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CsFetchHistoryMsgLogic) CsFetchHistoryMsg(in *pb.CsFetchHistoryMsgReq) (*pb.CsFetchHistoryMsgResp, error) {
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
		return nil, errors.Wrap(result.NewErrMsg("fetch cs history message list error"), "")
	}
	return &pb.CsFetchHistoryMsgResp{
		TotalPage:   1,
		CurrentPage: 1,
		List:        list,
	}, nil
}
