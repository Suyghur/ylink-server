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
	manager.flowMap.Insert(flow.FlowId, flow)
}

func (manager *flowManager) registerFlow(flow *model.Flow) {
	go manager.subscribeRmq(flow)
	for {
		select {
		case <-flow.Stream.Context().Done():
			if manager.Has(flow.FlowId) {
				flow.Logger.Infof("flowstream was disconnected abnormally")
				manager.UnRegister(flow.FlowId)
				flow.SvcCtx.InnerRpc.NotifyUserOffline(flow.Ctx, &inner.NotifyUserStatusReq{
					Type:   flow.Type,
					Uid:    flow.Uid,
					GameId: flow.GameId,
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
			resultCmd := flow.SvcCtx.RedisClient.BRPop(flow.Ctx, 10*time.Second, flow.FlowId)
			if message, err := resultCmd.Result(); err != nil {
				flow.Logger.Errorf("get message from redis, key: %s, err: %v", flow.FlowId, err)
			} else {
				flow.Message <- message[1]
			}
		}
	}
}

func (manager *flowManager) Get(flowId string) *model.Flow {
	return manager.flowMap.Get(flowId).(*model.Flow)
}

func (manager *flowManager) UnRegister(flowId string) {
	if manager.flowMap.Contains(flowId) {
		flow := manager.Get(flowId)
		close(flow.Message)
		//flow.EndRmq <- 0
		manager.flowMap.Erase(flowId)
	}
}

func (manager *flowManager) Has(flowId string) bool {
	return manager.flowMap.Contains(flowId)
}
