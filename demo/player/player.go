//@Author   : KaiShin
//@Time     : 2022/3/16

package main

import (
	"bufio"
	"call_center/call/rpc/call"
	"call_center/call/rpc/pb"
	"call_center/demo/data"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"io"
	"os"
)

func playerLogin() {
	conf := new(zrpc.RpcClientConf)
	conf.Target = data.Addr
	data.Client = call.NewCall(zrpc.MustNewClient(*conf))

	req := new(call.ClientMsgReq)
	cmd := pb.CommandMsg{}

	idInfo := pb.IdInfo{GameId: data.GameId} // 游客登录不需要id
	req.Cmd = append(req.Cmd, &cmd)
	req.IdInfo = &idInfo
	stream, err := data.Client.ClientLogin(context.Background(), req)
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
			case pb.ECommand_SEND_MSG:
				msg := cmd.GetChatMsg()
				logx.Infof("[客户端收到]：%s", msg.Input)
			case pb.ECommand_ON_PLAYER_CONNECT:
				data.PlayerId = cmd.CmdStr
				logx.Info("player conn, id: ", data.PlayerId)
				break
			case pb.ECommand_ON_SERVICE_DISCONNECT:
				logx.Infof("服务客服{%v}已断开", cmd.CmdStr)
				break
			case pb.ECommand_ON_PLAYER_RECEIVE_REPLY:
				data.ServiceId = cmd.CmdStr
				logx.Info("分配到客服：", data.ServiceId)
			default:
				break
			}
		}
	}
}

func playerOnCall() {
	input := bufio.NewReader(os.Stdin)
	for {
		logx.Info("请输入信息:")
		cmdStr, _ := input.ReadString('\n')

		var proto = new(pb.ClientMsgReq)

		var cmd = new(pb.CommandMsg)
		cmd.CmdType = pb.ECommand_CALL_PLAYER_MSG

		var chatMsg = new(pb.CommandMsg_ChatMsg)
		chatMsg.ChatMsg = &pb.ChatMsg{Input: cmdStr}
		cmd.Buff = chatMsg

		proto.Cmd = append(proto.Cmd, cmd)

		var idInfo = new(pb.IdInfo)
		idInfo.Id = data.PlayerId
		proto.IdInfo = idInfo

		if data.Client == nil {
			logx.Error("player not playerLogin, client is null")
			continue
		}

		if _, err := data.Client.ClientCall(context.Background(), proto); err != nil {
			logx.Error("call error:", err)
			continue
		}
	}
}

func main() {
	go playerOnCall()
	playerLogin()
}
