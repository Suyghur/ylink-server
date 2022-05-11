//@File     message.go
//@Time     2022/05/10
//@Author   #Suyghur,

package model

type ChatMessage struct {
	CreateTime string `json:"create_time"`
	Content    string `json:"content"`
	Pic        string `json:"pic"`
	ReceiverId string `json:"receiver_id"`
	SenderId   string `json:"sender_id"`
}
