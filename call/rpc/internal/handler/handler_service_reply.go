//@Author   : KaiShin
//@Time     : 2021/11/2

package handler

import (
	"call_center/call/rpc/internal/core"
	"call_center/call/rpc/internal/interaction"
	"call_center/call/rpc/pb"
	pb2 "call_center/db/rpc/pb"
	"call_center/public/exception"
	"fmt"
)

func ServiceReply(server *core.Server, serverId string, msg *pb.CommandMsg, interDb interaction.InterDb) error {
	var err error

	exception.Try(func() {
		toPlayerId := msg.CmdStr
		player := server.GetPlayer(toPlayerId)
		if player == nil {
			// 玩家不存在
			exception.Throw(fmt.Sprintf("<serviceReplyHandler> player not exist, id:%v", toPlayerId))
		}

		res := server.OnPlayerQuitWaitQueue(server, player)
		if res == nil {
			// 玩家不在等待队列中
			exception.Throw(fmt.Sprintf("<serviceReplyHandler> player not in wait queue, id:%v", toPlayerId))
		}

		// 建立连接
		err := server.OnConfirmConn(server, serverId, toPlayerId)
		if err != nil {
			exception.Throw(err)
		}

		// 记录日志
		interDb.ChatRecordToDb(player, serverId, toPlayerId, pb2.EDbRecordState_E_DB_RECORD_STATE_REPLY)
	}).Catch(func(e exception.Exception) {
		err = e.(error)
	}).Finally(func() {

	})

	return err
}
