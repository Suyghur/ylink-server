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
	case globalkey.ConnectTypeNormalPlayer:
		// 修改玩家在线状态
		if ext.GameOnlinePlayerMap.Contains(in.GameId) {
			// 有则取出玩家的map
			onlinePlayerMap := ext.GameOnlinePlayerMap.Get(in.GameId).(*treemap.Map)
			if onlinePlayerMap.Contains(in.Uid) {
				l.Logger.Error("该玩家已在线")
				// TODO 单点在线
			} else {
				ts := time.Now().Unix()
				playerInfo := &model.PlayerInfo{
					GameId:     in.GameId,
					PlayerId:   in.Uid,
					IsVip:      0,
					CsId:       "",
					ConnectTs:  ts,
					LastChatTs: 0,
					EnqueueTs:  ts,
					DequeueTs:  0,
				}
				//if playerInfo == nil {
				//	l.Logger.Infof("playerInfo is nil")
				//}
				onlinePlayerMap.Insert(in.Uid, playerInfo)
				// 放入等待队列
				ext.WaitingQueue.Insert(in.GameId+"_"+in.Uid, playerInfo)
				l.Logger.Infof("enqueue waiting list: %s", ext.WaitingQueue)
				//TODO 返回等待信息
			}
		} else {
			onlinePlayerMap := treemap.New(treemap.WithGoroutineSafe())
			ts := time.Now().Unix()
			playerInfo := &model.PlayerInfo{
				GameId:    in.GameId,
				PlayerId:  in.Uid,
				ConnectTs: ts,
				EnqueueTs: ts,
			}
			onlinePlayerMap.Insert(in.Uid, playerInfo)
			// 放入等待队列
			ext.WaitingQueue.Insert(in.GameId+"_"+in.Uid, playerInfo)
			l.Logger.Infof("enqueue waiting list: %s", ext.WaitingQueue)
			ext.GameOnlinePlayerMap.Insert(in.GameId, onlinePlayerMap)
			//TODO 返回等待信息

		}
	case globalkey.ConnectTypeVipPlayer:
		var onlinePlayerMap *treemap.Map
		if ext.GameOnlinePlayerMap.Contains(in.GameId) {
			onlinePlayerMap = ext.GameOnlinePlayerMap.Get(in.GameId).(*treemap.Map)
		} else {
			onlinePlayerMap = treemap.New(treemap.WithGoroutineSafe())
		}

		if playerInfo := ext.GetVipPlayer(in.GameId, in.Uid); playerInfo != nil {
			playerInfo.ConnectTs = time.Now().Unix()
			onlinePlayerMap.Insert(in.Uid, playerInfo)
			ext.GameOnlinePlayerMap.Insert(in.GameId, onlinePlayerMap)
		} else {
			return nil, errors.Wrap(result.NewErrMsg("用户不存在"), "")
		}
	case globalkey.ConnectTypeCs:
		if csInfo := ext.GetCsInfo(in.Uid); csInfo != nil {
			csInfo.OnlineStatus = 1
			//TODO 返回等待信息
		} else {
			return nil, errors.Wrap(result.NewErrMsg("用户不存在"), "")
		}
	default:
		return nil, errors.Wrap(result.NewErrMsg("用户不存在"), "")
	}
	return &pb.NotifyUserStatusResp{}, nil
}
