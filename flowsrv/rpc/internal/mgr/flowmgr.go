//@File     flowmgr.go
//@Time     2022/5/30
//@Author   #Suyghur,

package mgr

import (
	treemap "github.com/liyue201/gostl/ds/map"
	"sync"
	"time"
	"ylink/comm/result"
	"ylink/core/inner/rpc/inner"
	"ylink/flowsrv/rpc/internal/model"
	"ylink/flowsrv/rpc/pb"
)

type flowManager struct {
	flowMap *treemap.Map
}

var (
	instance *flowManager
	once     sync.Once
)

func GetFlowMgrInstance() *flowManager {
	once.Do(func() {
		instance = &flowManager{
			flowMap: treemap.New(treemap.WithGoroutineSafe()),
		}
	})
	return instance
}

func (manager *flowManager) Register(flow *model.Flow) {
	//go registerWorker(flow)
	go manager.registerFlow(flow)
	manager.flowMap.Insert(flow.User.Uid, flow)
}

func (manager *flowManager) registerFlow(flow *model.Flow) {
	go manager.subscribeRmq(flow)
	for {
		select {
		case <-flow.Stream.Context().Done():
			if manager.Has(flow.User.Uid) {
				flow.Logger.Infof("flowstream was disconnected abnormally")
				manager.UnRegister(flow.User.Uid)
				flow.SvcCtx.InnerRpc.NotifyUserOffline(flow.Ctx, &inner.NotifyUserStatusReq{
					Type:   flow.User.Type,
					Uid:    flow.User.Uid,
					GameId: flow.User.GameId,
				})
			}
			flow.EndFlow <- 1
			return
		case msg, open := <-flow.Message:
			if open {
				flow.Stream.Send(&pb.CommandResp{
					Code: result.Ok,
					Msg:  "success",
					Data: []byte(msg),
				})
			} else {
				flow.Logger.Error("message channel is close")
				return
			}
		}
	}
}

func (manager *flowManager) subscribeRmq(flow *model.Flow) {
	for {
		select {
		case <-flow.Stream.Context().Done():
			flow.Logger.Infof("unsubscribe rmq...")
			return
		default:
			resultCmd := flow.SvcCtx.RedisClient.BRPop(flow.Ctx, 10*time.Second, flow.User.Uid)
			if message, err := resultCmd.Result(); err != nil {
				flow.Logger.Errorf("get message from redis err: %v", err)
			} else {
				flow.Message <- message[1]
			}
		}
	}
}

func (manager *flowManager) Get(uid string) *model.Flow {
	return manager.flowMap.Get(uid).(*model.Flow)
}

func (manager *flowManager) UnRegister(uid string) {
	if manager.flowMap.Contains(uid) {
		flow := manager.Get(uid)
		close(flow.Message)
		//flow.EndRmq <- 0
		manager.flowMap.Erase(uid)
	}
}

func (manager *flowManager) Has(uid string) bool {
	return manager.flowMap.Contains(uid)
}
