//@Author   : KaiShin
//@Time     : 2021/10/28

package handler

import (
	"call_center/db/rpc/internal/svc"
	"call_center/db/rpc/pb"
	"fmt"
	"log"
)

func GetChatLog(svcCtx *svc.ServiceContext, cmdMsg *pb.DbCommandMsg) *pb.DbCommandMsg {
	playerId := cmdMsg.CmdStr

	querySql := fmt.Sprintf(`select * from service_chat_visitor_log_r where sessionId = "%s" limit 100`, playerId)
	log.Println("<handler.HandleGetChatLog> querySql:", querySql)

	dataRes := svcCtx.Es.Query(querySql)
	if dataRes == nil {
		return &pb.DbCommandMsg{}
	}

	chatLogList := new(pb.ArrayDbChatLog)
	for _, v := range dataRes {
		data := v.(map[string]interface{})
		chatLogList.DataList = append(chatLogList.DataList, &pb.DbChatLog{
			GameId:    int32(data["gameId"].(float64)),
			Content:   data["content"].(string),
			FromId:    data["fromId"].(string),
			ToId:      data["toId"].(string),
			ChatType:  pb.EDbChatType(data["chatType"].(float64)),
			SessionId: data["sessionId"].(string),
			// IsVisitor: data["isVisitor"].(bool),
			TimeStamp: int64(data["logDt"].(float64)),
		})
	}
	log.Println("<handler.GetChatLog>, data len: ", len(dataRes), " playerId:", playerId)
	return &pb.DbCommandMsg{Data: &pb.DbCommandMsg_ArrayChatLog{ArrayChatLog: chatLogList}}
}
