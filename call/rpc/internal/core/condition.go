package core

import (
	"call_center/call/rpc/internal/interaction"
	"call_center/call/rpc/internal/role"
	"call_center/call/rpc/pb"
	"errors"
	"log"
)

type Condition struct {
}

func (sel *Condition) OnPlayerConnect(server *Server, stream interface{}) (*role.Player, error) {
	/*
		玩家连接
	*/
	id := server.MakeId(&stream)

	cmd := server.QuickBuildCmdMsg(pb.ECommand_ON_PLAYER_CONNECT, 0, id)
	err := server.CmdToClient(stream, cmd)
	if err != nil {
		return nil, err
	}

	// record
	player := server.ConnPlayer(id, stream)

	// 推送当前排队信息
	server.WaitQueueLenUpdate([]interface{}{player})
	log.Println("<Condition.OnPlayerConnect> stream id:", id)
	return player, err
}

func (sel *Condition) OnPlayerDisConnect(server *Server, id string, reason pb.ErrorReason) {
	/*
		玩家断开
	*/

	/*
		通知相关
	*/
	cmd := server.QuickBuildCmdMsg(pb.ECommand_ON_PLAYER_DISCONNECT, int32(reason), id)

	// 通知玩家离线
	player := server.GetPlayer(id)
	pStream := player.Stream
	if pStream != nil {
		err := server.CmdToClient(pStream, cmd)
		if err != nil {
			log.Printf("<Condition.OnPlayerDisConnect> server.CmdToClient err:%s, id:%s", err, id)
		}
	}

	//如果有对应客服，通知客服玩家离线
	service := server.GetServiceByPlayerId(id)
	if service != nil && reason != pb.ErrorReason_SERVICE_HEART_BEAT_FAILED {
		err := server.CmdToService(service.Stream, cmd)
		if err != nil {
			log.Printf("<Condition.OnPlayerDisConnect> server.GetService err:%s, id:%s", err, id)
		}

		// 推送玩家挂断列表
		idInfo := pb.IdInfo{GameId: player.GameId, Id: player.Id}
		service.HangUpList = append(service.HangUpList, &idInfo)
		server.PushHangUpList(service)
	}

	// 清除玩家记录
	server.DisConnPlayer(id)
	log.Printf("<Condition.OnPlayerDisConnect> End, id: %s, reason:%d", id, reason)
}

func (sel *Condition) OnServiceConnect(serviceId string, server *Server, stream interface{}, db interaction.InterDb) (*role.Service, error) {
	/*
		客服连接
	*/
	id := server.MakeId(&stream)
	if serviceId != "" {
		id = serviceId
	}

	cmd := server.QuickBuildCmdMsg(pb.ECommand_ON_SERVICE_CONNECT, 0, id)
	err := server.CmdToService(stream, cmd)
	if err != nil {
		return nil, err
	}

	// 日志record
	service := server.ConnService(id, stream)

	// 玩家等待队列通知
	server.WaitQueueInfoUpdate(server.waitQueue.GetAll(), []interface{}{service})

	// 挂断列表历史记录
	recordList := db.GetChatRecord(id)
	if recordList != nil {
		service.InitHandUpList(recordList)
		server.PushHangUpList(service)
	}

	log.Println("<Condition.OnServiceConnect> id:", id)
	return service, nil
}

func (sel *Condition) OnServiceDisConnect(server *Server, id string, errCode pb.ErrorReason) {
	/*
		客服断开
	*/

	// 踢出对接中的玩家
	{
		pidList := server.GetPlayersByServiceId(id)
		for _, pid := range pidList {
			server.KickPlayer(pid.(string), int32(errCode))
		}
	}

	connPIds := server.DisConnService(id)
	for _, pcId := range connPIds {
		stream := server.GetPlayerStream(pcId.(string))
		if stream == nil {
			continue
		}
		cmd := server.QuickBuildCmdMsg(pb.ECommand_ON_SERVICE_DISCONNECT, int32(errCode), id)
		err := server.CmdToClient(stream, cmd)
		if err != nil {
			continue
		}
	}
	log.Println("<Condition.OnServiceDisConnect> id: ", id)
}

func (sel *Condition) OnPlayerEnterWaitQueue(server *Server, player *role.Player) {
	/*
		玩家加入等待队列
	*/
	server.AddWaitQueue(player)

	// 更新队列信息to客服
	waitPlayers := server.waitQueue.GetAll()
	noticeService := server.GetAllService()
	server.WaitQueueInfoUpdate(waitPlayers, noticeService)

	// 更新排队人数to玩家
	queueLen := server.waitQueue.Len()
	server.WaitQueueLenUpdate(waitPlayers)

	log.Printf("<Condition.OnPlayerEnterWaitQueue> id: %s, queueLen:%d", player.Id, queueLen)
}

func (sel *Condition) OnPlayerQuitWaitQueue(server *Server, player *role.Player) interface{} {
	/*
		玩家移除等待队列
	*/
	res := server.RemoveWaitQueue(player)
	if res != nil {
		// 更新队列信息to客服
		waitPlayers := server.waitQueue.GetAll()
		noticeService := server.GetAllService()
		server.WaitQueueInfoUpdate(waitPlayers, noticeService)

		// 更新当前排队人数to玩家
		server.WaitQueueLenUpdate(waitPlayers)
	}
	log.Printf("<Condition.OnPlayerQuitWaitQueue> id:%s, queueLen:%d ", player.Id, server.waitQueue.Len())
	return res
}

func (sel *Condition) OnConfirmConn(server *Server, sId string, pId string) error {
	/*
		确认分配玩家到对应客服
	*/

	// 建立连接
	ok := server.ConfirmService(sId, pId)
	if ok != true {
		return errors.New("ConfirmService failed")
	}

	// 通知客服
	var cmd = new(pb.CommandMsg)
	cmd.CmdType = pb.ECommand_ON_PLAYER_RECEIVE_REPLY
	cmd.CmdStr = pId

	stream := server.GetServiceStream(sId)
	err := server.CmdToService(stream, cmd)
	if err != nil {
		log.Println("<Condition.OnConfirmConn> CmdToService err:", err)
	}

	// 通知玩家
	cmd.CmdStr = sId
	stream = server.GetPlayerStream(pId)
	err = server.CmdToClient(stream, cmd)
	if err != nil {
		log.Println("<Condition.OnConfirmConn> CmdToClient err:", err)
	}

	return nil
}
