package logic

import (
	"context"
	"github.com/liyue201/gostl/ds/set"
	"github.com/pkg/errors"
	"ylink/comm/globalkey"
	"ylink/comm/model"
	"ylink/comm/result"
	"ylink/core/inner/rpc/internal/ext"

	"ylink/core/inner/rpc/internal/svc"
	"ylink/core/inner/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotifyUserOfflineLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNotifyUserOfflineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotifyUserOfflineLogic {
	return &NotifyUserOfflineLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *NotifyUserOfflineLogic) NotifyUserOffline(in *pb.NotifyUserStatusReq) (*pb.NotifyUserStatusResp, error) {
	switch in.Type {
	case globalkey.CONNECT_TYPE_PLAYER:
		// 修改玩家在线状态
		if ext.Game2PlayerStatMap.Contains(in.GameId) {
			// 有则取出玩家的set
			playerStatSet := ext.Game2PlayerStatMap.Get(in.GameId).(*set.Set)
			if playerStatSet.Contains(in.Uid) {
				// 有则清除，代表下线
				playerStatSet.Erase(in.Uid)
			}
		}

		for n := ext.WaitingQueue.FrontNode(); n != nil; n = n.Next() {
			info := n.Value.(*model.PlayerWaitingInfo)
			if info.GameId == in.GameId && info.PlayerId == in.Uid {
				l.Logger.Infof("remove the player from the queue, game_id: %s, player_id: %s", in.GameId, in.Uid)
				ext.WaitingQueue.Remove(nil, n)
				break
			}
		}
		l.Logger.Infof("waiting queue size: %d", ext.WaitingQueue.Len())
		l.Logger.Infof("waiting queue: %s", ext.WaitingQueue.String())
	case globalkey.CONNECT_TYPE_CS:
		// 修改客服在线状态
		csInfo := ext.CsMap.Get(in.Uid).(*model.CsInfo)
		csInfo.OnlineStatus = 0
	default:
		return nil, errors.Wrap(result.NewErrMsg("no such user type"), "")
	}
	return &pb.NotifyUserStatusResp{}, nil
}
