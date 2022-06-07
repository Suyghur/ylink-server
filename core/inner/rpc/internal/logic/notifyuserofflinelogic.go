package logic

import (
	"context"
	treemap "github.com/liyue201/gostl/ds/map"
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
		if ext.GameOnlinePlayerMap.Contains(in.GameId) {
			// 有则取出玩家
			onlinePlayerMap := ext.GameOnlinePlayerMap.Get(in.GameId).(*treemap.Map)
			if onlinePlayerMap.Contains(in.Uid) {
				// 有则清除，代表下线
				onlinePlayerMap.Erase(in.Uid)
			}
		}

		for n := ext.WaitingList.FrontNode(); n != nil; n = n.Next() {
			info := n.Value.(*model.PlayerInfo)
			if info.GameId == in.GameId && info.PlayerId == in.Uid {
				l.Logger.Infof("remove the player from the queue, game_id: %s, player_id: %s", in.GameId, in.Uid)
				ext.WaitingList.Remove(nil, n)
				break
			}
		}
	case globalkey.CONNECT_TYPE_CS:
		// 修改客服在线状态
		if csInfo := ext.GetCsInfo(in.Uid); csInfo != nil {
			csInfo.OnlineStatus = 0
		} else {
			return nil, errors.Wrap(result.NewErrMsg("no such user"), "")
		}
	default:
		return nil, errors.Wrap(result.NewErrMsg("no such user type"), "")
	}
	return &pb.NotifyUserStatusResp{}, nil
}
