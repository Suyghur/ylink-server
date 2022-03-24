package handler

import (
	"call_center/call/rpc/internal/core"
	"call_center/call/rpc/internal/interaction"
	"call_center/call/rpc/pb"
	"errors"
)

func FromPlayerMsg(interDb interaction.InterDb, server *core.Server, streamId string, msg *pb.CommandMsg) error {
	chatMsg := msg.GetChatMsg()
	player := server.GetPlayer(streamId)
	input := chatMsg.Input
	player.RefreshTalkTimeStamp() // 刷新发言时间

	// 接收数据
	// gid := common.GetGoroutineID()
	// log.Printf("[gid: %v][收到消息]： %s from: %v \n", gid, input, streamId)

	service := server.GetServiceByPlayerId(streamId)
	if service != nil {
		// 人工接听时， 消息返回给客服
		err := server.MsgToService(service.Stream, input, streamId)
		if err != nil {
			return err
		}
		interDb.ChatLogToDb(player, streamId, service.Id, input, true)
		// err = server.MsgToClient(stream, "success")
	} else {
		return errors.New("<FromPlayerMsg> 客服不存在")
		//// 机器自动回复
		//robotMsg := "[robot回复]" + input
		//err := server.MsgToClient(stream, robotMsg)
		//// interaction.DbInst.ChatLogToDb(streamId, "", input, true)
		//
		//if err != nil {
		//	return err
		//}

	}

	return nil
}
