package logic

import (
	"context"
	treemap "github.com/liyue201/gostl/ds/map"
	"github.com/liyue201/gostl/ds/set"
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
		if ext.Game2PlayerStatusMap.Contains(in.GameId) {
			// 有则取出玩家的set
			playerStatSet := ext.Game2PlayerStatusMap.Get(in.GameId).(*set.Set)
			if !playerStatSet.Contains(in.Uid) {
				playerStatSet.Insert(in.Uid)
			}
		} else {
			playerStatSet := set.New()
			playerStatSet.Insert(in.Uid)
			ext.Game2PlayerStatusMap.Insert(in.GameId, playerStatSet)
		}

		// 判断是否有专属客服，没有则放入等待队列
		if ext.GameVipMap.Contains(in.GameId) {
			p2cMap := ext.GameVipMap.Get(in.GameId).(*treemap.Map)
			if !p2cMap.Contains(in.Uid) {
				ext.WaitingQueue.PushBack(&model.PlayerWaitingInfo{
					PlayerId:    in.Uid,
					GameId:      in.GameId,
					EnqueueTime: time.Now().Unix(),
				})
			}
		} else {
			ext.WaitingQueue.PushBack(&model.PlayerWaitingInfo{
				PlayerId:    in.Uid,
				GameId:      in.GameId,
				EnqueueTime: time.Now().Unix(),
			})
		}
		l.Logger.Infof("enqueue waiting list: %s", ext.WaitingQueue.String())
	case globalkey.CONNECT_TYPE_CS:
		// 修改客服在线状态
		csInfo := ext.CsInfoMap.Get(in.Uid).(*model.CsInfo)
		csInfo.OnlineStatus = 1
	default:
		return nil, errors.Wrap(result.NewErrMsg("no such user type"), "")
	}
	return &pb.NotifyUserStatusResp{}, nil
}
