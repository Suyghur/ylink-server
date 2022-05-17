//@File     message.go
//@Time     2022/05/10
//@Author   #Suyghur,

package model

type KqMessage struct {
	CreateTime  string `json:"create_time"`
	Content     string `json:"content"`
	Pic         string `json:"pic"`
	ReceiverId  string `json:"receiver_id"`
	SenderId    string `json:"sender_id"`
	GameId      string `json:"game_id"`
	OperationId string `json:"operation_id"`
}

type KqCmdMessage struct {
	KqMessage
	Opt int64 `json:"opt"`
}
