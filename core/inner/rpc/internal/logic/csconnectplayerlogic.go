package logic

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/gookit/event"
	treemap "github.com/liyue201/gostl/ds/map"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
	"ylink/comm/globalkey"
	"ylink/comm/model"
	"ylink/comm/result"
	"ylink/core/inner/rpc/internal/ext"
	"ylink/core/inner/rpc/internal/svc"
	"ylink/core/inner/rpc/pb"
)

type CsConnectPlayerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCsConnectPlayerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsConnectPlayerLogic {
	return &CsConnectPlayerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CsConnectPlayerLogic) CsConnectPlayer(in *pb.InnerCsConnectPlayerReq) (*pb.InnerCsConnectPlayerResp, error) {

	playerInfo := ext.GetOnlinePlayerInfo(in.GameId, in.PlayerId)

	if playerInfo == nil {
		return nil, errors.Wrapf(result.NewErrMsg("The player is not connected"), "")
	}
	playerInfo.CsId = in.CsId
	playerInfo.DequeueTs = time.Now().Unix()

	if ext.GameConnectedMap.Contains(in.GameId) {
		playerConnectedMap := ext.GameConnectedMap.Get(in.GameId).(*treemap.Map)
		playerConnectedMap.Insert(in.PlayerId, playerInfo)
	} else {
		playerConnectedMap := treemap.New(treemap.WithGoroutineSafe())
		playerConnectedMap.Insert(in.PlayerId, playerInfo)
		ext.GameConnectedMap.Insert(in.GameId, playerConnectedMap)
	}

	// 移除WaitingQueue
	for n := ext.WaitingList.FrontNode(); n != nil; n = n.Next() {
		playerInfo := n.Value.(*model.PlayerInfo)
		if playerInfo.GameId == in.GameId && playerInfo.PlayerId == in.PlayerId {
			l.Logger.Infof("remove the player from the queue, game_id: %s, player_id: %s", in.GameId, in.PlayerId)
			ext.WaitingList.Remove(nil, n)
			break
		}
	}

	var entryId cron.EntryID
	entryId, _ = l.svcCtx.TimeoutCron.AddFunc("@every 1m", func() {
		// TODO 增加trace
		var timeoutTs int64
		if playerInfo.LastChatTs == 0 {
			timeoutTs = time.Now().Unix() - playerInfo.ConnectTs
		} else {
			timeoutTs = time.Now().Unix() - playerInfo.LastChatTs
		}
		if timeoutTs >= 3600 {
			// 释放计时器任务
			_ = event.MustFire(globalkey.EventRemoveTimeoutJob, event.M{"entry_id": entryId})
			l.Logger.Infof("trigger timeout event, remove cron job, entry id: %d", entryId)

			// 发踢下线的command指令
			message, _ := sonic.MarshalString(&model.KqCmdMessage{
				Opt:        model.CMD_CHAT_TIMEOUT,
				ReceiverId: in.GameId + "_" + in.PlayerId,
				GameId:     in.GameId,
				Uid:        in.PlayerId,
			})
			l.svcCtx.KqCmdBoxProducer.SendMessage(l.ctx, message, in.GameId+"_"+in.PlayerId)
		}
	})

	return &pb.InnerCsConnectPlayerResp{}, nil
}
