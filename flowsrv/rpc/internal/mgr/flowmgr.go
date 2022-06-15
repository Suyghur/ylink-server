//@File     flowmgr.go
//@Time     2022/5/30
//@Author   #Suyghur,

package mgr

import (
	"context"
	treemap "github.com/liyue201/gostl/ds/map"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
	"time"
	"ylink/comm/ctxdata"
	"ylink/comm/result"
	"ylink/comm/trace"
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
				flow.Logger.Infof("stream was disconnected abnormally")
				manager.UnRegister(flow.FlowId)
				manager.handleUserOffline(flow)
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
	traceId := ctxdata.GetTraceIdFromCtx(flow.Stream.Context())
	trace.RunOnTracing(traceId, func(ctx context.Context) {
		for {
			select {
			case <-flow.Stream.Context().Done():
				logx.WithContext(ctx).Infof("unsubscribe rmq...")
				return
			default:
				resultCmd := flow.SvcCtx.RedisClient.BRPop(ctx, 30*time.Second, flow.FlowId)
				if message, err := resultCmd.Result(); err != nil {
					logx.WithContext(ctx).Errorf("get message from redis, key: %s, err: %v", flow.FlowId, err)
				} else {
					trace.StartTrace(ctx, "FlowsrvServer.flowmgr.handleRmqMessage", func(ctx context.Context) {
						flow.Message <- message[1]
					})
				}
			}
		}
	})
}

func (manager *flowManager) handleUserOffline(flow *model.Flow) {
	traceId := ctxdata.GetTraceIdFromCtx(flow.Stream.Context())
	trace.RunOnTracing(traceId, func(ctx context.Context) {
		trace.StartTrace(ctx, "FlowsrvServer.flowmgr.handleUserOffline", func(ctx context.Context) {
			_, err := flow.SvcCtx.InnerRpc.NotifyUserOffline(ctx, &inner.NotifyUserStatusReq{
				Type:   flow.Type,
				Uid:    flow.Uid,
				GameId: flow.GameId,
			})
			if err != nil {
				logx.WithContext(ctx).Errorf("notify user offline has some error: %v", err)
			}
		})
	})
}

func (manager *flowManager) Get(flowId string) *model.Flow {
	return manager.flowMap.Get(flowId).(*model.Flow)
}

func (manager *flowManager) UnRegister(flowId string) {
	if manager.flowMap.Contains(flowId) {
		flow := manager.Get(flowId)
		close(flow.Message)
		manager.flowMap.Erase(flowId)
	}
}

func (manager *flowManager) Has(flowId string) bool {
	return manager.flowMap.Contains(flowId)
}
