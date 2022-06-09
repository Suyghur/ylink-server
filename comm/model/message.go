//@File     message.go
//@Time     2022/05/10
//@Author   #Suyghur,

package model

const (
	CMD_SEND_MESSAGE = 0
	CMD_CHAT_TIMEOUT = 2001
)

type KqMessage struct {
	CreateTime string `json:"create_time"`
	Content    string `json:"content"`
	Pic        string `json:"pic"`
	ReceiverId string `json:"receiver_id"`
	SenderId   string `json:"sender_id"`
	GameId     string `json:"game_id"`
	Uid        string `json:"uid"`
	Ext        string `json:"ext"`
}

type KqCmdMessage struct {
	Opt        int64  `json:"opt"`
	ReceiverId string `json:"receiver_id"`
	GameId     string `json:"game_id"`
	Uid        string `json:"uid"`
	Ext        string `json:"ext"`
}
