// Code generated by goctl. DO NOT EDIT.
package types

type CommResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PlayerAuthReq struct {
	PlayerId string `json:"player_id"`
	GameId   string `json:"game_id"`
}

type PlayerFetchCsInfoReq struct {
	CsId string `json:"cs_id"`
}

type PlayerFetchHistoryMsgReq struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type PlayerSendMsgReq struct {
	Content string `json:"content"`
	Pic     string `json:"pic"`
}

type CsAuthReq struct {
	UserName string `json:"uname"`
	Password string `json:"password"`
}

type CsFetchPlayerInfoReq struct {
	PlayerId string `json:"palyer_id"`
	GameId   string `json:"game_id"`
}

type CsFetchPlayerQueueReq struct {
	Limit int `json:"limit"`
}

type CsConnectPlayerReq struct {
	PlayerId string `json:"player_id"`
	GameId   string `json:"game_id"`
}

type CsFetchHistoryChatReq struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type CsFetchHistoryMsgReq struct {
	PlayerId string `json:"player_id"`
	GameId   string `json:"game_id"`
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
}

type CsFetchMsgReq struct {
	PlayerId string `json:"player_id"`
	GameId   string `json:"game_id"`
}

type CsSendMsgReq struct {
	PlayerId string `json:"player_id"`
	GameId   string `json:"game_id"`
	Content  string `json:"content"`
	Pic      string `json:"pic"`
}
