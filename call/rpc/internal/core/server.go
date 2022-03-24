package core

import (
	"call_center/call/rpc/internal/role"
	"call_center/call/rpc/pb"
	public "call_center/public/common"
	"container/list"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"log"
	"time"
)

var instance *Server

func ServerInit(conf Config) {
	instance = new(Server)

	// 加载配置
	instance.Config = conf

	// 等待队列
	instance.waitQueue = public.NewQueue()

	log.Println("<Server.ServerInit> end, conf:", conf)
}

func GetServer() *Server {
	return instance
}

type Server struct {
	Condition
	Communication
	Config

	// 玩家流管理类
	playerMgr public.ObjMgr // pId -> Player

	// 客服流管理类
	serviceMgr public.ObjMgr // sId -> Service

	// 玩家 -> 客服 映射管理类
	p2sMgr public.ObjMgr // pId -> sId

	// 客服 -> 玩家 映射管理类
	s2pMgr public.ObjMgr // sId -> set(pId)

	// 等待队列
	waitQueue *public.SyncQueue
}

func (sel *Server) MakeId(i interface{}) string {
	now := time.Now().Unix()
	objId := fmt.Sprintf("%d%d", now, i)
	return objId
}

func (sel *Server) GetServiceStream(scId string) interface{} {
	service := sel.serviceMgr.GetObj(scId)
	if service != nil {
		return service.(*role.Service).Stream
	}
	return nil
}

func (sel *Server) GetPlayerStream(pcId string) interface{} {
	player := sel.playerMgr.GetObj(pcId)
	if player != nil {
		return player.(*role.Player).Stream
	}
	return nil
}

func (sel *Server) GetPlayer(pId string) *role.Player {
	p := sel.playerMgr.GetObj(pId)
	if p != nil {
		return p.(*role.Player)
	}
	return nil
}

func (sel *Server) GetService(sId string) *role.Service {
	s := sel.serviceMgr.GetObj(sId)
	if s != nil {
		return s.(*role.Service)
	}

	return nil
}

func (sel *Server) GetServiceByPlayerId(pId string) *role.Service {
	scId := sel.p2sMgr.GetObj(pId)
	if scId == nil {
		return nil
	}
	return sel.GetService(scId.(string))
}

func (sel *Server) GetPlayersByServiceId(sid string) []interface{} {
	var res []interface{}
	setPid := sel.s2pMgr.GetObj(sid)
	if setPid != nil {
		ss := setPid.(mapset.Set)
		return ss.ToSlice()
	}
	return res
}

func (sel *Server) GetAllService() []interface{} {
	return sel.serviceMgr.GetObjValues()
}

func (sel *Server) ConfirmService(scId string, pcId string) bool {
	curScId := sel.p2sMgr.GetObj(pcId)
	if curScId != nil {
		if curScId == scId {
			// 当前客服对应玩家信息一致
			return true
		}
		log.Printf("<Server.ConfirmService> player<%v> already in service by %v \n", pcId, curScId)
		return false
	}

	sel.p2sMgr.Register(pcId, scId)
	obj := sel.s2pMgr.GetObj(scId)
	idSet := obj.(mapset.Set)
	idSet.Add(pcId)
	log.Printf("<Server.ConfirmService> pcId:{%v}, scId:{%v}, idSet:{%v} \n", pcId, scId, idSet.String())
	return true
}

func (sel *Server) ConnPlayer(pcId string, stream interface{}) *role.Player {
	// 玩家连接

	/*
		注册玩家信息
	*/
	pInfo := new(role.Player)
	pInfo.Id = pcId
	pInfo.Stream = stream

	sel.playerMgr.Register(pcId, pInfo)
	return pInfo
}

func (sel *Server) KickPlayer(pId string, reason int32) {
	log.Printf("<Server.KickPlayer>, id:%s, reason:%d", pId, reason)
	player := sel.playerMgr.GetObj(pId)
	if player != nil {
		p := player.(*role.Player)
		p.StopChan(reason)

		log.Println("<Server.KickPlayer> <- begin kick player, id:", pId)
		<-p.WaitLogOut()
		log.Println("<Server.KickPlayer> -> end kick player, id:", pId)
	}
}

func (sel *Server) KickService(sId string, reason int32) {
	log.Printf("<Server.KickService>, id:%s, reason:%d", sId, reason)
	service := sel.serviceMgr.GetObj(sId)
	if service != nil {
		s := service.(*role.Service)
		s.StopChan(reason)

		log.Println("<Server.KickService> <- begin kick service, id:", sId)
		<-s.WaitLogOut()
		log.Println("<Server.KickService> -> end kick service, id:", sId)
	}
}

func (sel *Server) DisConnPlayer(pcId string) {
	// 玩家连接关闭
	player := sel.playerMgr.GetObj(pcId)

	/*
		删除wait queue
	*/
	sel.waitQueue.Remove(player)

	/*
		删除玩家info
	*/

	sel.playerMgr.DeleteObj(pcId)

	/*
		删除对接该玩家的客服映射
	*/
	scId := sel.p2sMgr.GetObj(pcId)
	if scId != nil {
		obj := sel.s2pMgr.GetObj(scId)
		if obj != nil {
			idSet := obj.(mapset.Set)
			idSet.Remove(pcId)
			log.Printf("<Server.DisConnPlayer> disConnPcId: {%v} curScId: {%v} -> pcIds: {%v} \n", pcId, scId, idSet.String())
		}
	}
	sel.p2sMgr.DeleteObj(pcId)
	player.(*role.Player).Final()
	log.Printf("<Server.DisConnPlayer> rmPcId: %v, scId: %v", pcId, scId)
}

func (sel *Server) ConnService(scId string, stream interface{}) *role.Service {
	// 客服连接

	/*
		注册客服stream
	*/
	service := new(role.Service)
	service.Id = scId
	service.Stream = stream

	sel.serviceMgr.Register(scId, service)

	/*
	* 创建客服映射
	 */
	idSet := mapset.NewSet()
	sel.s2pMgr.Register(scId, idSet)

	return service
}

func (sel *Server) DisConnService(scId string) []interface{} {
	// 客服连接关闭
	service := sel.serviceMgr.GetObj(scId)
	sel.serviceMgr.DeleteObj(scId)

	/*
		删除该客服对接的所有玩家
	*/
	// 遍历字典，获取该scId所服务的pcId
	var delList []interface{}
	obj := sel.s2pMgr.GetObj(scId)
	if obj == nil {
		return delList
	}
	idSet := obj.(mapset.Set)
	it := idSet.Iterator()
	for itId := range it.C {
		sel.p2sMgr.DeleteObj(itId)
		delList = append(delList, itId)
	}
	sel.s2pMgr.DeleteObj(scId)
	service.(*role.Service).Final()
	log.Printf("<Server.DisConnService> scId: %v, selId: %v \n", scId, delList)
	return delList
}

func (sel *Server) AddWaitQueue(player *role.Player) {
	queueLen := sel.waitQueue.PushBack(player)
	time.AfterFunc(time.Second*time.Duration(sel.WaitConnServiceLimit), sel.verifyWaitQueue) // 初始化一个校验等待队列计时器
	log.Printf("<Server.AddWaitQueue> playerId:%s, wait sec:%d, queue_len:%d \n",
		player.Id, sel.WaitConnServiceLimit, queueLen)
}

func (sel *Server) RemoveWaitQueue(player *role.Player) interface{} {
	res := sel.waitQueue.Remove(player)
	log.Printf("<Server.RemoveWaitQueue> playerId:%s, queue_len:%d \n", player.Id, sel.waitQueue.Len())
	return res
}

func (sel *Server) verifyWaitQueue() {
	log.Println("<Server.verifyWaitQueue> begin, queue len:", sel.waitQueue.Len())
	back := sel.waitQueue.Back()
	if back == nil {
		return
	}
	e := back.(*list.Element)
	player := e.Value.(*role.Player)
	nowTimeStamp := time.Now().Unix()
	offset := nowTimeStamp - player.LoginTimeStamp
	if offset >= sel.WaitConnServiceLimit {
		log.Println("<Server.verifyWaitQueue> wait overtime, kick out id:", player.Id)
		// sel.waitQueue.RemoveE(e)
		sel.KickPlayer(player.Id, int32(pb.ErrorReason_PLAYER_WAIT_QUEUE_OVERTIME))
	}
	log.Println("<Server.verifyWaitQueue> end, queue len:", sel.waitQueue.Len())
}
