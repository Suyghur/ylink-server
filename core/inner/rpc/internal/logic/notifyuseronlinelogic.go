package logic

import (
	"context"
	treemap "github.com/liyue201/gostl/ds/map"
	"github.com/pkg/errors"
	"time"
	"ylink/comm/globalkey"
	"ylink/comm/model"
	"ylink/comm/result"
	"ylink/core/inner/rpc/internal/ext"

	"ylink/core/inner/rpc/internal/svc"
	"ylink/core/inner/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotifyUserOnlineLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNotifyUserOnlineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotifyUserOnlineLogic {
	return &NotifyUserOnlineLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *NotifyUserOnlineLogic) NotifyUserOnline(in *pb.NotifyUserStatusReq) (*pb.NotifyUserStatusResp, error) {
	switch in.Type {
	case globalkey.CONNECT_TYPE_PLAYER:
		// 修改玩家在线状态
		if ext.GameOnlinePlayerMap.Contains(in.GameId) {
			// 有则取出玩家的map
			onlinePlayerMap := ext.GameOnlinePlayerMap.Get(in.GameId).(*treemap.Map)
			if onlinePlayerMap.Contains(in.Uid) {
				l.Logger.Error("such player has been connected")
			} else {
				// 不存在换这个玩家，判断是否vip
				if playerInfo := ext.GetVipPlayer(in.GameId, in.Uid); playerInfo != nil {
					playerInfo.ConnectTs = time.Now().Unix()
					onlinePlayerMap.Insert(in.Uid, playerInfo)
				} else {
					// 不是vip
					onlinePlayerMap.Insert(in.Uid, &model.PlayerInfo{
						PlayerId:  in.Uid,
						GameId:    in.GameId,
						ConnectTs: time.Now().Unix(),
					})
					// 放入等待队列
					ext.WaitingQueue.PushBack(&model.PlayerWaitingInfo{
						PlayerId:    in.Uid,
						GameId:      in.GameId,
						EnqueueTime: time.Now().Unix(),
					})
					l.Logger.Infof("enqueue waiting list: %s", ext.WaitingQueue.String())
				}
			}
		} else {
			onlinePlayerMap := treemap.New(treemap.WithGoroutineSafe())
			// 判断是不是vip玩家
			if playerInfo := ext.GetVipPlayer(in.GameId, in.Uid); playerInfo != nil {
				playerInfo.ConnectTs = time.Now().Unix()
				onlinePlayerMap.Insert(in.Uid, playerInfo)
			} else {
				// 不是vip
				onlinePlayerMap.Insert(in.Uid, &model.PlayerInfo{
					PlayerId:  in.Uid,
					GameId:    in.GameId,
					ConnectTs: time.Now().Unix(),
				})
				// 放入等待队列
				ext.WaitingQueue.PushBack(&model.PlayerWaitingInfo{
					PlayerId:    in.Uid,
					GameId:      in.GameId,
					EnqueueTime: time.Now().Unix(),
				})
				l.Logger.Infof("enqueue waiting list: %s", ext.WaitingQueue.String())
			}
			ext.GameOnlinePlayerMap.Insert(in.GameId, onlinePlayerMap)
		}
	case globalkey.CONNECT_TYPE_CS:
		if csInfo := ext.GetCsInfo(in.Uid); csInfo != nil {
			csInfo.OnlineStatus = 1
		} else {
			return nil, errors.Wrap(result.NewErrMsg("no such user"), "")
		}
	default:
		return nil, errors.Wrap(result.NewErrMsg("no such user type"), "")
	}
	return &pb.NotifyUserStatusResp{}, nil
}
