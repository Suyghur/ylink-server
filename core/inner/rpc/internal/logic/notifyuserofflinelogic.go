package logic

import (
	"context"
	treemap "github.com/liyue201/gostl/ds/map"
	"github.com/pkg/errors"
	"ylink/comm/globalkey"
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
	l.Logger.Infof("NotifyUserOffline")
	switch in.Type {
	case globalkey.ConnectTypeNormalPlayer:
		// 修改玩家在线状态
		if ext.GameOnlinePlayerMap.Contains(in.GameId) {
			// 有则取出玩家
			onlinePlayerMap := ext.GameOnlinePlayerMap.Get(in.GameId).(*treemap.Map)
			if onlinePlayerMap.Contains(in.Uid) {
				// 有则清除，代表下线
				onlinePlayerMap.Erase(in.Uid)
				l.Logger.Infof("清除玩家在线状态")
			}
		}

		key := in.GameId + "_" + in.Uid
		if ext.WaitingQueue.Contains(key) {
			l.Logger.Infof("remove the player from the queue, game_id: %s, player_id: %s", in.GameId, in.Uid)
			ext.WaitingQueue.Erase(key)

		}
		//for n := ext.WaitingList.FrontNode(); n != nil; n = n.Next() {
		//	info := n.Value.(*model.PlayerInfo)
		//	l.Logger.Infof("playerInfo: %v", info)
		//	if info.GameId == in.GameId && info.PlayerId == in.Uid {
		//		l.Logger.Infof("remove the player from the queue, game_id: %s, player_id: %s", in.GameId, in.Uid)
		//		ext.WaitingList.Remove(nil, n)
		//		//TODO 广播客户端更新等待队列信息
		//		//queueInfo, _ := sonic.MarshalString(map[string]interface{}{
		//		//	"queue_size": 10,
		//		//})
		//		//message, _ := sonic.MarshalString(&model.KqCmdMessage{
		//		//	Opt:        model.CMD_UPDATE_WAITING_QUEUE,
		//		//	ReceiverId: "all",
		//		//	GameId:     in.GameId,
		//		//	Uid:        in.Uid,
		//		//	Ext:        queueInfo,
		//		//})
		//		break
		//	}
		//}
	case globalkey.ConnectTypeVipPlayer:
		// 修改玩家在线状态
		if ext.GameOnlinePlayerMap.Contains(in.GameId) {
			// 有则取出玩家
			onlinePlayerMap := ext.GameOnlinePlayerMap.Get(in.GameId).(*treemap.Map)
			if onlinePlayerMap.Contains(in.Uid) {
				// 有则清除，代表下线
				onlinePlayerMap.Erase(in.Uid)
				l.Logger.Infof("清除玩家在线状态")
			}
		}
	case globalkey.ConnectTypeCs:
		// 修改客服在线状态
		if csInfo := ext.GetCsInfo(in.Uid); csInfo != nil {
			csInfo.OnlineStatus = 0
		} else {
			return nil, errors.Wrap(result.NewErrMsg("用户不存在"), "")
		}
	default:
		return nil, errors.Wrap(result.NewErrMsg("用户不存在"), "")
	}
	return &pb.NotifyUserStatusResp{}, nil
}
