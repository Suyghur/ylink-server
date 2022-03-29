package logic

import (
	"call_center/call/rpc/internal/core"
	"call_center/call/rpc/pb"
	db "call_center/db/rpc/pb"
	"call_center/public/exception"
	"context"
	"log"
	"time"

	"call_center/call/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClientLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClientLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientLoginLogic {
	return &ClientLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClientLoginLogic) ClientLogin(req *pb.ClientMsgReq, stream pb.Call_ClientLoginServer) error {
	server := core.GetServer()

	idInfo := req.IdInfo
	if idInfo == nil {
		// 传参失效
		errStr := "<ClientLogin> req.IdInfo is nil"
		return exception.MakeError(int32(pb.EErrorCode_ERR_PARAM_ERROR), errStr)
	}

	if idInfo.Id != "" {
		// 重登
		p := server.GetPlayer(idInfo.Id)

		if p != nil {
			// 当前已连接,断开当前链接
			server.KickPlayer(idInfo.Id, int32(pb.ErrorReason_PLAYER_REPEAT_LOGIN))
			log.Println("<ClientLogin> player already conn, disconnect first, id:", idInfo.Id)
		}
	}

	// 客服stream注册
	player, err := server.OnPlayerConnect(server, stream)
	if player == nil {
		log.Println("<ClientLogin> OnPlayerConnect failed, err:", err)
		//return err
		return exception.MakeError(int32(pb.EErrorCode_ERR_PLAYER_CONN_ERR), err.Error())
	}

	// 初始化玩家信息
	stopChan := make(chan int32)
	player.Init(idInfo.GameId, true, stopChan)

	// 加入等待队列
	server.OnPlayerEnterWaitQueue(server, player)

	// 心跳ticker
	duration := time.Second * time.Duration(server.HeartBeatInterval)
	ticker := time.NewTicker(duration)
	var errCode pb.ErrorReason
	exception.Try(func() {
		defer func() {
			ticker.Stop()
		}()
		for {
			select {
			case <-stream.Context().Done():
				log.Println("<ClientLogin> heartbeat failed, id:", player.Id, " err:", stream.Context().Err())
				errCode = pb.ErrorReason_PLAYER_HEART_BEAT_FAILED
				return
			case <-ticker.C:
				// 超时未发言
				now := time.Now().Unix()
				lastTalkTimeStamp := player.LastTalkTimeStamp
				if now-lastTalkTimeStamp >= server.LastTalkIntervalLimit {
					log.Println("<ClientLogin> last talk interval limit, id:", player.Id)
					errCode = pb.ErrorReason_PLAYER_TALK_INTERVAL_LIMIT
					return
				}
			case stop := <-stopChan:
				// 登出信号
				errCode = pb.ErrorReason(stop)
				log.Println("<ClientLogin> player stop connect, code:", stop)
				return
			}
		}
	}).Catch(func(ex exception.Exception) {
		log.Println("<ClientLogin> error:", ex)
	}).Finally(func() {
		switch errCode {
		// 处理等待队列中的玩家
		case pb.ErrorReason_PLAYER_WAIT_QUEUE_OVERTIME, pb.ErrorReason_PLAYER_CALL_LOGOUT, pb.ErrorReason_PLAYER_REPEAT_LOGIN:
			server.OnPlayerQuitWaitQueue(server, player)
			break
		}

		// 记录日志
		service := server.GetServiceByPlayerId(player.Id)
		if service != nil {
			l.svcCtx.Db.ChatRecordToDb(player, service.Id, player.Id, db.EDbRecordState_E_DB_RECORD_STATE_DISCONNECT)
		}

		// 断开处理
		server.OnPlayerDisConnect(server, player.Id, errCode)
	})
	return nil
}
