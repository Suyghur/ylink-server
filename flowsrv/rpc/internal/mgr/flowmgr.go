//@File     flowmgr.go
//@Time     2022/05/13
//@Author   #Suyghur,

package mgr

import (
	"sync"
	"ylink/flowsrv/rpc/pb"
)

type flowManager struct {
	flowMap map[string]pb.Flowsrv_ConnectServer
}

var (
	instance *flowManager
	once     sync.Once
)

func GetFlowMgrInstance() *flowManager {
	once.Do(func() {
		instance = &flowManager{
			flowMap: make(map[string]pb.Flowsrv_ConnectServer),
		}
	})
	return instance
}

func (manager *flowManager) SetFlow(uid string, flow pb.Flowsrv_ConnectServer) {
	manager.flowMap[uid] = flow
}

func (manager *flowManager) GetFlow(uid string) pb.Flowsrv_ConnectServer {
	return manager.flowMap[uid]
}

func (manager *flowManager) RemoveFlow(uid string) {
	delete(manager.flowMap, uid)
}
