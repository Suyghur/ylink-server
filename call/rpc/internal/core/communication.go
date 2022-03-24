package core

import (
	"call_center/call/rpc/internal/role"
	"call_center/call/rpc/pb"
	"log"
)

type Communication struct {
}

func (sel *Communication) _protoToClient(stream interface{}, proto *pb.ClientMsgRes) error {
	var err error
	ss := stream.(pb.Call_ClientLoginServer)
	err = ss.Send(proto)

	if err != nil {
		log.Println("<_protoToClient> error: ", err)
		return err
	}

	return nil
}

func (sel *Communication) _protoToService(stream interface{}, proto *pb.ServiceMsgRes) error {
	var err error
	ss := stream.(pb.Call_ServiceLoginServer)
	err = ss.Send(proto)

	if err != nil {
		log.Println("<_protoToService> error: ", err)
		return err
	}

	return nil
}

func (sel *Communication) QuickBuildCmdMsg(cmdType interface{}, val int32, str string) *pb.CommandMsg {
	var cmd = new(pb.CommandMsg)
	cmd.CmdType = cmdType.(pb.ECommand)
	cmd.CmdVal = val
	cmd.CmdStr = str
	return cmd
}

func (sel *Communication) CmdToService(stream interface{}, cmd *pb.CommandMsg) error {
	var res = new(pb.ServiceMsgRes)
	res.Cmd = append(res.Cmd, cmd)
	if cmd.CmdType != pb.ECommand_MSG_HEART_BEAT {
		log.Println("[DEBUG]<CmdToService> type:", cmd.CmdType, " cmd:", cmd)
	}
	return sel._protoToService(stream, res)
}

func (sel *Communication) CmdToClient(stream interface{}, cmd *pb.CommandMsg) error {
	var res = new(pb.ClientMsgRes)
	res.Cmd = append(res.Cmd, cmd)
	if cmd.CmdType != pb.ECommand_MSG_HEART_BEAT {
		log.Println("[DEBUG]<CmdToClient> type:", cmd)
	}
	return sel._protoToClient(stream, res)
}

func (sel *Communication) CmdBroadcastService(arrayService []interface{}, cmd *pb.CommandMsg) {
	for _, pInst := range arrayService {
		service := pInst.(*role.Service)
		stream := service.Stream
		err := sel.CmdToService(stream, cmd)
		if err != nil {
			log.Println("<Communication.CmdBroadcastService> send failed, id:", service.Id, "error:", err)
			continue
		}
	}
}

func (sel *Communication) CmdBroadcastClient(arrayPlayer []interface{}, cmd *pb.CommandMsg) {
	for _, pInst := range arrayPlayer {
		player := pInst.(*role.Player)
		stream := player.Stream
		err := sel.CmdToClient(stream, cmd)
		if err != nil {
			log.Println("<Communication.MsgBroadcastClient> send failed, id:", player.Id, "error:", err)
			continue
		}
	}
}

func (sel *Communication) MsgToClient(stream interface{}, content string, sendId string) error {
	var proto = new(pb.ClientMsgRes)

	var cmd = new(pb.CommandMsg)
	cmd.CmdType = pb.ECommand_SEND_MSG
	cmd.Buff = &pb.CommandMsg_ChatMsg{ChatMsg: &pb.ChatMsg{Input: content, ClientId: sendId}}
	proto.Cmd = append(proto.Cmd, cmd)

	return sel._protoToClient(stream, proto)
}

func (sel *Communication) MsgToService(stream interface{}, content string, sendId string) error {
	var proto = new(pb.ServiceMsgRes)

	var cmd = new(pb.CommandMsg)
	cmd.CmdType = pb.ECommand_SEND_MSG
	cmd.Buff = &pb.CommandMsg_ChatMsg{ChatMsg: &pb.ChatMsg{Input: content, ClientId: sendId}}
	proto.Cmd = append(proto.Cmd, cmd)
	return sel._protoToService(stream, proto)
}

func (sel *Communication) ServiceHeartBeat(stream interface{}) error {
	cmd := sel.QuickBuildCmdMsg(pb.ECommand_MSG_HEART_BEAT, 0, "")
	return sel.CmdToService(stream, cmd)
}

func (sel *Communication) PlayerHeartBeat(player *role.Player) error {
	var err error
	stream := player.Stream
	cmd := sel.QuickBuildCmdMsg(pb.ECommand_MSG_HEART_BEAT, 0, "")
	err = sel.CmdToClient(stream, cmd)
	if err != nil {
		return err
	}
	return err
}

func (sel *Communication) WaitQueueInfoUpdate(waitPlayers, noticedService []interface{}) {
	if len(noticedService) == 0 {
		return
	}
	cmdMsg := new(pb.CommandMsg)
	arrayIdInfo := pb.CommandMsg_ArrayIdInfo{ArrayIdInfo: &pb.ArrayIdInfo{}}
	for _, pInst := range waitPlayers {
		player := pInst.(*role.Player)
		arrayIdInfo.ArrayIdInfo.IdInfos = append(arrayIdInfo.ArrayIdInfo.IdInfos, &pb.IdInfo{Id: player.Id, GameId: player.GameId})
	}
	cmdMsg.Buff = &arrayIdInfo
	cmdMsg.CmdType = pb.ECommand_ON_PLAYER_WAIT_QUEUE_INFO
	sel.CmdBroadcastService(noticedService, cmdMsg)
}

func (sel *Communication) WaitQueueLenUpdate(waitPlayers []interface{}) {
	if len(waitPlayers) == 0 {
		return
	}
	cmdMsg := new(pb.CommandMsg)
	cmdMsg.CmdType = pb.ECommand_ON_PLAYER_WAIT_QUEUE_LEN

	for idx, pInst := range waitPlayers {
		cmdMsg.CmdVal = int32(idx)
		player := pInst.(*role.Player)
		stream := player.Stream
		err := sel.CmdToClient(stream, cmdMsg)
		if err != nil {
			log.Println("<Communication.WaitQueueLenUpdate> error: ", err)
			continue
		}
	}
}

func (sel *Communication) PushHangUpList(service *role.Service) {
	hangUpCmd := sel.QuickBuildCmdMsg(pb.ECommand_ON_SERVICE_HANG_UP_LIST, 0, "")
	arrayIdInfo := pb.CommandMsg_ArrayIdInfo{ArrayIdInfo: &pb.ArrayIdInfo{}}
	arrayIdInfo.ArrayIdInfo.IdInfos = service.HangUpList
	hangUpCmd.Buff = &arrayIdInfo
	err := sel.CmdToService(service.Stream, hangUpCmd)
	if err != nil {
		log.Println("<Communication.PushHangUpList>, err:", err, " serviceId:", service.Id)
	}
}
