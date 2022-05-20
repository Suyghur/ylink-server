//@File     playerinfo.go
//@Time     2022/05/19
//@Author   #Suyghur,

package model

type PlayerWaitingInfo struct {
	PlayerId    string `json:"player_id"`
	GameId      string `json:"game_id"`
	EnqueueTime int64  `json:"enqueue_time"`
}
