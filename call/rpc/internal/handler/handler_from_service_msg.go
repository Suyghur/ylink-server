//@Author   : KaiShin
//@Time     : 2021/11/2

package handler

import (
	"call_center/call/rpc/internal/core"
	"call_center/call/rpc/internal/interaction"
	"call_center/call/rpc/pb"
	"call_center/public/exception"
	"fmt"
	"log"
)

func FromServiceMsg(server *core.Server, streamId string, msg *pb.CommandMsg, interDb interaction.InterDb) error {
	// 客服连接请求
	chatMsg := msg.GetChatMsg()
	clientId := chatMsg.ClientId
	player := server.GetPlayer(clientId)
	// log.Println("Recv, 客服回复玩家<", clientId, ">：", chatMsg.Input)

	var err error = nil
	exception.Try(func() {
		if player == nil {
			// 玩家不存在
			errStr := fmt.Sprintf("<serviceMsgHandler> 玩家id{%v}不存在", clientId)
			exception.Throw(errStr)
		}

		// 客服回复玩家
		err = server.MsgToClient(player.Stream, chatMsg.Input, streamId)
		if err != nil {
			log.Println("<serviceMsgHandler> err:", err)
			exception.Throw(err)
		}

		// 推送日志
		interDb.ChatLogToDb(player, clientId, streamId, chatMsg.Input, false)

		//// 发送成功确认信息
		//err = server.MsgToService(serviceStream, "success")
		//if err != nil {
		//	exception.Throw(err)
		//}
	}).Catch(func(e exception.Exception) {
		log.Println("<serviceMsgHandler> err:", e)
		err = e.(error)
	}).Finally(func() {

	})

	return err
}
