//@Author   : KaiShin
//@Time     : 2022/3/16

package main

import (
	"bufio"
	"call_center/call/rpc/call"
	"call_center/call/rpc/pb"
	"call_center/demo/data"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"io"
	"log"
	"os"
	"strings"
)

func serviceLogin() {
	conf := new(zrpc.RpcClientConf)
	conf.Target = data.Addr
	data.Client = call.NewCall(zrpc.MustNewClient(*conf))

	req := new(call.ServiceMsgReq)
	cmd := pb.CommandMsg{}

	idInfo := pb.IdInfo{GameId: data.GameId, Id: data.ServiceId}
	req.Cmd = append(req.Cmd, &cmd)
	req.IdInfo = &idInfo
	stream, err := data.Client.ServiceLogin(context.Background(), req)
	if err != nil {
		logx.Errorf("playerLogin failed, playerId:%s, err:%s", err)
		return
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			logx.Info("收到服务端的结束信号")
			break
		}

		if err != nil {
			logx.Info("接收数据错误：", err)
		}

		if res == nil {
			break
		}

		for _, cmd := range res.Cmd {
			if cmd.CmdType != pb.ECommand_MSG_HEART_BEAT {
				logx.Infof("[DEBUG] 收到cmd：%v, val:%v, str:%v, buff:%s", cmd.CmdType, cmd.CmdVal, cmd.CmdStr, cmd.GetBuff())
			}

			switch cmd.CmdType {
			case pb.ECommand_ON_PLAYER_RECEIVE_REPLY:
				data.PlayerId = cmd.CmdStr
				logx.Info("分配到玩家：", data.PlayerId)
			case pb.ECommand_SEND_MSG:
				buff := cmd.GetChatMsg()
				logx.Info("[客戶端收到]：", buff.Input)
			case pb.ECommand_ON_PLAYER_DISCONNECT:
				if data.ServiceId == cmd.CmdStr {
					data.PlayerId = "default"
					logx.Infof("玩家[%v]已断开\n", cmd.CmdStr)
				}
			case pb.ECommand_ON_SERVICE_CONNECT:
				data.ServiceId = cmd.CmdStr
			case pb.ECommand_ON_PLAYER_CONNECT:
				// 玩家登录
				if data.PlayerId == "default" {
					data.PlayerId = cmd.CmdStr
					logx.Infof("玩家{%v}登录，客服{%v}当前空闲，已分配", data.PlayerId, data.ServiceId)
				}
			default:
				break
			}
		}
	}
}

func serviceOnCall() {
	input := bufio.NewReader(os.Stdin)

	for {
		logx.Info("请输入信息:\n")
		chatContent, _ := input.ReadString('\n')

		if data.Client == nil {
			logx.Error("service not login, client is null")
			continue
		}

		var proto = new(pb.ServiceMsgReq)
		var cmd = new(pb.CommandMsg)
		cmd = parseGmCmd(chatContent)
		if cmd != nil {

		} else {
			cmd = new(pb.CommandMsg)
			cmd.CmdType = pb.ECommand_CALL_SERVICE_MSG
			var chatMsg = new(pb.CommandMsg_ChatMsg)
			chatMsg.ChatMsg = &pb.ChatMsg{ClientId: data.PlayerId, Input: chatContent}
			cmd.Buff = chatMsg
		}

		proto.Cmd = append(proto.Cmd, cmd)
		proto.IdInfo = &pb.IdInfo{Id: data.ServiceId}

		if res, err := data.Client.ServiceCall(context.Background(), proto); err != nil {
			fmt.Println("error:", err)
			continue
		} else {
			if res != nil {
				log.Println(res)
			}
		}
	}
}

func parseGmCmd(cmd string) *pb.CommandMsg {
	/*
		 客服连接
		{"cmd_type": 2005, "cmd_val": 1, "cmd_str": "1648121824824636975904"}

		获取聊天日志
		{"cmd_type": 2006, "cmd_val": 2, "cmd_str": "16474122581374391221744"}
	*/

	ok := strings.Contains(cmd, "cmd_type")
	if ok != true {
		return nil
	}
	cmd = strings.Replace(cmd, "\n", "", -1)

	cmdMsg := new(pb.CommandMsg)

	var cmdStr map[string]interface{}
	err := json.Unmarshal([]byte(cmd), &cmdStr)
	if err != nil {
		log.Println("<ParseCmd> err:", err)
		return nil
	}

	cmdTpe := int32(cmdStr["cmd_type"].(float64))
	cmdVal := int32(cmdStr["cmd_val"].(float64))
	cmdMsg.CmdType = pb.ECommand(cmdTpe)
	cmdMsg.CmdVal = cmdVal
	cmdMsg.CmdStr = cmdStr["cmd_str"].(string)
	return cmdMsg
}

func main() {
	go serviceLogin()
	serviceOnCall()
}
