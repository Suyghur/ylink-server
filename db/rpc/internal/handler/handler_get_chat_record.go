//@Author   : KaiShin
//@Time     : 2021/10/28

package handler

import (
	"call_center/db/rpc/internal/svc"
	"call_center/db/rpc/pb"
	"fmt"
	"log"
	"time"
)

func GetChatRecord(svcCtx *svc.ServiceContext, cmdMsg *pb.DbCommandMsg) *pb.DbCommandMsg {
	serviceId := cmdMsg.CmdStr

	// 只获取最近一天的记录
	now := time.Now().Unix()
	oneDayBefore := now - (24 * 3600)

	// 构建sql
	querySql := fmt.Sprintf(`select playerId, gameId from service_chat_record where state = 1 and serviceId = '%s' and logDt >= %d limit 100`, serviceId, oneDayBefore)
	log.Println("<handler.GetChatRecord> querySql:", querySql)

	dataRes := svcCtx.Es.Query(querySql)
	if dataRes == nil {
		return &pb.DbCommandMsg{}
	}

	chatRecordList := new(pb.ArrayChatRecord)
	for _, v := range dataRes {
		data := v.(map[string]interface{})
		chatRecordList.DataList = append(chatRecordList.DataList, &pb.DbChatRecord{
			PlayerId: data["playerId"].(string),
			GameId:   int32(data["gameId"].(float64)),
		})
	}
	return &pb.DbCommandMsg{Data: &pb.DbCommandMsg_ArrayChatRecord{ArrayChatRecord: chatRecordList}}
}
