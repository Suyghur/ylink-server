package logic

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/structpb"
	"ylink/core/inner/rpc/internal/svc"
	"ylink/core/inner/rpc/pb"
	"ylink/ext/result"

	"github.com/zeromicro/go-zero/core/logx"
)

type CsFetchPlayerQueueLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCsFetchPlayerQueueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsFetchPlayerQueueLogic {
	return &CsFetchPlayerQueueLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CsFetchPlayerQueueLogic) CsFetchPlayerQueue(in *pb.InnerCsFetchPlayerQueueReq) (*pb.InnerCsFetchPlayerQueueResp, error) {
	// todo: 查询等待用户的队列

	list, err := structpb.NewList([]interface{}{
		map[string]interface{}{
			"player_id": "player1111",
			"game_id":   "game1231",
			"wait_time": 1000,
		},
		map[string]interface{}{
			"player_id": "player2222",
			"game_id":   "game1231",
			"wait_time": 10,
		},
	})
	if err != nil {
		return nil, errors.Wrap(result.NewErrMsg("fetch player wait queue error"), "")
	}

	return &pb.InnerCsFetchPlayerQueueResp{
		List: list,
	}, nil
}
