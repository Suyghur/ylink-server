//@File     playerinfo.go
//@Time     2022/05/19
//@Author   #Suyghur,

package model

type PlayerInfo struct {
	GameId     string `json:"game_id"`
	PlayerId   string `json:"player_id"`
	IsVip      int32  `json:"is_vip"`
	CsId       string `json:"cs_id"`
	ConnectTs  int64  `json:"connect_ts"`
	LastChatTs int64  `json:"last_chat_ts"`
	EnqueueTs  int64  `json:"enqueue_ts"`
	DequeueTs  int64  `json:"dequeue_ts"`
}
