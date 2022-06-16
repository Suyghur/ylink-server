//@File     message.go
//@Time     2022/05/10
//@Author   #Suyghur,

package model

const (
	CMD_SEND_MESSAGE         = 0
	CMD_UPDATE_WAITING_QUEUE = 2000
	CMD_CHAT_TIMEOUT         = 2001
)

type KqMessage struct {
	Opt        int32  `json:"opt"`
	CreateTs   int64  `json:"create_ts"`
	Payload    string `json:"payload"`
	ReceiverId string `json:"receiver_id"`
	SenderId   string `json:"sender_id"`
	GameId     string `json:"game_id"`
	Uid        string `json:"uid"`
	Ext        string `json:"ext"`
}
type ChatMessage struct {
	CreateTime string `json:"create_time"`
	Content    string `json:"content"`
	Pic        string `json:"pic"`
}

type CommandMessage struct {
	CmdInfo interface{} `json:"cmd_info"`
}
