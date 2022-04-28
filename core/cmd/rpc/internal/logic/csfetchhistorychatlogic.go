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

type CsFetchHistoryChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCsFetchHistoryChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsFetchHistoryChatLogic {
	return &CsFetchHistoryChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CsFetchHistoryChatLogic) CsFetchHistoryChat(in *pb.CsFetchHistoryChatReq) (*pb.CsFetchHistoryChatResp, error) {
	list, err := structpb.NewList([]interface{}{
		map[string]interface{}{
			"player_id":         "test1231",
			"player_name":       "一条大菜狗",
			"player_avatar_url": "https://www.baidu.com",
			"game_id":           "game1231",
			"game_name":         "青云诀2",
			"update_time":       "2022-04-27 17:01:40",
		},
		map[string]interface{}{
			"player_id":         "test1111",
			"player_name":       "高手",
			"player_avatar_url": "https://www.baidu.com",
			"game_id":           "game1231",
			"game_name":         "青云诀2",
			"update_time":       "2022-04-27 17:01:40",
		},
	})
	if err != nil {
		return nil, errors.Wrap(result.NewErrMsg("fetch cs chat list error"), "")
	}
	return &pb.CsFetchHistoryChatResp{
		TotalPage:   1,
		CurrentPage: 1,
		List:        list,
	}, nil
}
