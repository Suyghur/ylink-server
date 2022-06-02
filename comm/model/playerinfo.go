//@File     playerinfo.go
//@Time     2022/05/19
//@Author   #Suyghur,

package model

type PlayerInfo struct {
	PlayerId   string `json:"player_id"`
	GameId     string `json:"game_id"`
	IsVip      int64  `json:"is_vip"`
	CsId       string `json:"cs_id"`
	ConnectTs  int64  `json:"connect_ts"`
	LastChatTs int64  `json:"last_chat_ts"`
}

type PlayerWaitingInfo struct {
	PlayerId    string `json:"player_id"`
	GameId      string `json:"game_id"`
	EnqueueTime int64  `json:"enqueue_time"`
}
