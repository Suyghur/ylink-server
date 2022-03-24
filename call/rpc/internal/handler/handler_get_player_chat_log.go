//@Author   : KaiShin
//@Time     : 2021/11/2

package handler

import (
	"call_center/call/rpc/internal/interaction"
	"call_center/call/rpc/pb"
	db "call_center/db/rpc/pb"
	"call_center/public/exception"
	"log"
)

func GetPlayerChatLog(serviceId string, req *pb.CommandMsg, interDb interaction.InterDb) (*pb.CommandMsg, error) {
	cmdMsg := new(pb.CommandMsg)
	exception.Try(func() {
		playerId := req.CmdStr
		dataList := interDb.GetChatLog(playerId)
		if dataList != nil {
			chatLogList := new(pb.ArrayChatLog)
			for _, v := range dataList {
				fromPlayer := false
				if v.ChatType == db.EDbChatType_E_DB_CHAT_TYPE_PLAYER {
					fromPlayer = true
				}
				chatLogList.DataList = append(chatLogList.DataList, &pb.ChatLog{
					Content:    v.Content,
					TimeStamp:  v.TimeStamp,
					GameId:     v.GameId,
					FromPlayer: fromPlayer,
				})
			}
			cmdMsg.Buff = &pb.CommandMsg_ArrayChatLog{ArrayChatLog: chatLogList}
		}
		log.Println("<handler.playerChatLogHandler> id:", serviceId)
	}).Catch(func(e exception.Exception) {

	}).Finally(func() {

	})
	return cmdMsg, nil
}
