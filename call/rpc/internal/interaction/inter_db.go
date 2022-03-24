package interaction

import (
	"call_center/call/rpc/internal/role"
	"call_center/db/rpc/pb"
	"call_center/public/exception"
	"context"
	"encoding/json"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

/*
	数据库交互类
*/

type (
	InterDb interface {
		Start()
		ChatLogToDb(player *role.Player, playerId string, serviceId string, content string, fromPlayer bool)
		ChatRecordToDb(player *role.Player, serviceId string, playerId string, state pb.EDbRecordState)
		GetChatRecord(serviceId string) []*pb.DbChatRecord
		GetChatLog(playerId string) []*pb.DbChatLog
	}

	interDb struct {
		client   pb.DbClient
		dbStream pb.Db_DbLoginClient
	}
)

func NewInterDb(cli pb.DbClient) InterDb {
	log.Println("<NewInterDb> cli:", cli)
	dbInst := &interDb{client: cli}
	return dbInst
}

func (sel *interDb) Start() {
	log.Println("<Start> client:", sel.client)
	duration := time.Second * 5
	ticker := time.NewTicker(duration)
	for range ticker.C {
		exception.Try(func() {
			sel.runDb()
		}).Catch(func(e exception.Exception) {

		}).Finally(func() {

		})
	}
}

func (sel *interDb) runDb() {
	ctx := context.Background()

	// 赋值client
	dbClient := sel.client

	// context塞入信息
	md := metadata.Pairs("Key", "Val")
	ctx = metadata.NewOutgoingContext(ctx, md)
	stream, err := dbClient.DbLogin(ctx)
	if err != nil {
		// log.Println(err)
		return
	}

	log.Printf("<inter_db.runDB> db连接成功， 地址：%v", &dbClient)

	// 赋值stream
	sel.dbStream = stream

	// 获取配置信息
	req := new(pb.DbMsgReq)
	cmd := new(pb.DbCommandMsg)
	cmd.CmdType = pb.EDbCommand_E_DB_COMMAND_GET_CONFIG
	req.Cmd = cmd
	if res, err := dbClient.DbCall(ctx, req); err == nil {
		arrList := res.GetCmd().GetArrayConfig().GetDataList()
		for _, conf := range arrList {
			var confValueMap map[string]interface{}
			confName := conf.ConfName
			confKey := conf.ConfKey
			err := json.Unmarshal([]byte(conf.ConfValue), &confValueMap)
			if err != nil {
				continue
			}
			log.Println(confName, confKey, confValueMap)
		}
		log.Println(arrList)
	}

	for {
		receive, err := stream.Recv()
		if err != nil {
			sel.dbStream = nil
			log.Println("<inter_db.runDB> conn broken, err:", err)
			break
		}
		//res := receive.GetCmd()
		//switch res.CmdType {
		//
		//}
		log.Println("<inter_db.runDB> begin receive...:", receive)
	}

	log.Println("<inter_db.runDB> dbClient stopped, begin to reconnect...")
}

func (sel *interDb) send(cmd *pb.DbCommandMsg) {
	if sel.dbStream == nil {
		log.Println("<interDb.send> dbStream is nil")
		return
	}

	err := sel.dbStream.Send(&pb.DbMsgReq{Cmd: cmd})
	if err != nil {
		log.Println("<interDb.send> Send err:", err)
		return
	}
}

func (sel *interDb) ChatLogToDb(player *role.Player, playerId string, serviceId string, content string, fromPlayer bool) {
	if sel.dbStream == nil {
		log.Println("<interDb.ChatLogToDb> InterDb not init")
	}
	exception.Try(func() {
		chatInfo := new(pb.DbChatLog)
		chatInfo.Content = content
		chatInfo.TimeStamp = time.Now().Unix()
		chatInfo.SessionId = player.SessionId
		chatInfo.IsVisitor = player.IsVisitor
		chatInfo.GameId = player.GameId
		if fromPlayer == true {
			chatInfo.ChatType = pb.EDbChatType_E_DB_CHAT_TYPE_PLAYER
			chatInfo.FromId = playerId
			chatInfo.ToId = serviceId
		} else {
			chatInfo.ChatType = pb.EDbChatType_E_DB_CHAT_TYPE_SERVICE
			chatInfo.FromId = serviceId
			chatInfo.ToId = playerId
		}

		cmd := new(pb.DbCommandMsg)
		cmd.CmdType = pb.EDbCommand_E_DB_COMMAND_PUSH_CHAT_LOG
		cmd.Data = &pb.DbCommandMsg_ChatLog{ChatLog: chatInfo}

		sel.send(cmd)
		log.Println("<interDb.ChatLogToDb> msg: ", chatInfo)

	}).Catch(func(e exception.Exception) {

	}).Finally(func() {

	})
}

func (sel *interDb) ChatRecordToDb(player *role.Player, serviceId string, playerId string, state pb.EDbRecordState) {
	if sel.dbStream == nil {
		log.Println("<interDb.ChatRecordToDb> InterDb not init")
	}
	exception.Try(func() {
		recordInfo := new(pb.DbChatRecord)
		recordInfo.TimeStamp = time.Now().Unix()
		recordInfo.SessionId = player.SessionId
		recordInfo.IsVisitor = player.IsVisitor
		recordInfo.GameId = player.GameId
		recordInfo.State = state
		recordInfo.PlayerId = playerId
		recordInfo.ServiceId = serviceId

		cmd := new(pb.DbCommandMsg)
		cmd.CmdType = pb.EDbCommand_E_DB_COMMAND_PUSH_CHAT_RECORD
		cmd.Data = &pb.DbCommandMsg_ChatRecord{ChatRecord: recordInfo}

		sel.send(cmd)
		log.Println("<interDb.ChatRecordToDb> msg: ", recordInfo)

	}).Catch(func(e exception.Exception) {

	}).Finally(func() {

	})
}

func (sel *interDb) GetChatRecord(serviceId string) []*pb.DbChatRecord {
	if sel.dbStream == nil {
		log.Println("<interDb.ChatLogToDb> InterDb not init")
	}

	req := new(pb.DbMsgReq)
	req.Cmd = &pb.DbCommandMsg{CmdStr: serviceId, CmdType: pb.EDbCommand_E_DB_COMMAND_GET_CHAT_RECORD}

	res, err := sel.client.DbCall(context.Background(), req)
	if err != nil {
		log.Println("<interDb.GetChatRecord> err:", err, " id:", serviceId)
		return nil
	}

	resCmd := res.GetCmd()
	if resCmd != nil {
		arrRecord := resCmd.GetArrayChatRecord()
		if arrRecord != nil {
			return arrRecord.GetDataList()
		}
	}

	log.Println("<interDb.GetChatRecord> record empty, id:", serviceId)
	return nil
}

func (sel *interDb) GetChatLog(playerId string) []*pb.DbChatLog {
	req := new(pb.DbMsgReq)
	req.Cmd = &pb.DbCommandMsg{CmdStr: playerId, CmdType: pb.EDbCommand_E_DB_COMMAND_GET_CHAT_LOG}

	res, err := sel.client.DbCall(context.Background(), req)
	if err != nil {
		log.Println("<interDb.GetChatLog> err:", err, " id:", playerId)
		return nil
	}
	resCmd := res.GetCmd()
	if resCmd != nil {
		dataList := res.GetCmd().GetArrayChatLog()
		if dataList != nil {
			return dataList.GetDataList()
		}
	}
	return nil
}
