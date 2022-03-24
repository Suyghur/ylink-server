//@Author   : KaiShin
//@Time     : 2021/10/28

package handler

import (
	"call_center/db/rpc/internal/svc"
	"call_center/db/rpc/pb"
	"log"
)

func PushChatLog(svcCtx *svc.ServiceContext, cmdMsg *pb.DbCommandMsg) {
	logInfo := cmdMsg.GetChatLog()

	mapInfo := make(map[string]interface{})
	mapInfo["content"] = logInfo.GetContent()
	mapInfo["fromId"] = logInfo.GetFromId()
	mapInfo["toId"] = logInfo.GetToId()
	mapInfo["chatType"] = logInfo.GetChatType()
	mapInfo["logDt"] = logInfo.GetTimeStamp()
	mapInfo["gameId"] = logInfo.GetGameId()
	mapInfo["sessionId"] = logInfo.GetSessionId()

	insertTable := "service_chat_log"
	if logInfo.GetIsVisitor() == true {
		insertTable = "service_chat_visitor_log"
	}
	svcCtx.Es.Insert(insertTable, mapInfo)
	log.Printf("<handler.PushChatLog>, table:%s, msg:%s", insertTable, mapInfo)
}
