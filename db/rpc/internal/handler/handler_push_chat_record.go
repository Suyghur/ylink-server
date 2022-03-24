//@Author   : KaiShin
//@Time     : 2021/10/28

package handler

import (
	"call_center/db/rpc/internal/svc"
	"call_center/db/rpc/pb"
	"log"
)

func PushChatRecord(svcCtx *svc.ServiceContext, cmdMsg *pb.DbCommandMsg) {
	recordInfo := cmdMsg.GetChatRecord()

	var mapInfo map[string]interface{}
	mapInfo = make(map[string]interface{})
	mapInfo["serviceId"] = recordInfo.GetServiceId()
	mapInfo["playerId"] = recordInfo.GetPlayerId()
	mapInfo["sessionId"] = recordInfo.GetSessionId()
	mapInfo["isVisitor"] = recordInfo.GetIsVisitor()
	mapInfo["state"] = int32(recordInfo.GetState())
	mapInfo["logDt"] = recordInfo.GetTimeStamp()
	mapInfo["gameId"] = recordInfo.GetGameId()

	insertTable := "service_chat_record"
	svcCtx.Es.Insert(insertTable, mapInfo)
	log.Println("<handler.PushChatRecord>, msg:", mapInfo)
}
