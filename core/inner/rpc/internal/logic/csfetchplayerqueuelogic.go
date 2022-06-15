package logic

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/structpb"
	"time"
	"ylink/comm/model"
	"ylink/comm/result"
	"ylink/core/inner/rpc/internal/ext"
	"ylink/core/inner/rpc/internal/svc"
	"ylink/core/inner/rpc/pb"

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
	queueLen := int32(ext.WaitingList.Len())
	if queueLen == 0 {
		// 等待队列为空直接返回
		return &pb.InnerCsFetchPlayerQueueResp{
			List: nil,
		}, nil
	}

	var index int32 = 0
	if in.Limit != 0 && in.Limit < queueLen {
		queueLen = in.Limit
	}

	queue := make([]interface{}, queueLen)

	for node := ext.WaitingList.FrontNode(); node != nil && index < queueLen; node = node.Next() {
		info := node.Value.(*model.PlayerInfo)
		queue[index] = map[string]interface{}{
			"game_id":   info.GameId,
			"player_id": info.PlayerId,
			"wait_time": time.Now().Unix() - info.EnqueueTs,
		}
		index += 1
	}

	list, err := structpb.NewList(queue)
	if err != nil {
		return nil, errors.Wrap(result.NewErrMsg("fetch player wait queue error"), "")
	}

	return &pb.InnerCsFetchPlayerQueueResp{
		List: list,
	}, nil
}
