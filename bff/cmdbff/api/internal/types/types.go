// Code generated by goctl. DO NOT EDIT.
package types

type PlayerFetchCsInfoReq struct {
	CsId string `json:"cs_id"`
}

type PlayerFetchCsInfoResp struct {
	CsId         string `json:"cs_id"`
	CsNickname   string `json:"cs_nickname"`
	CsAvatarUrl  string `json:"cs_avatar_url"`
	CsSignature  string `json:"cs_signature"`
	OnlineStatus int32  `json:"online_status"`
}

type PlayerFetchHistoryMsgReq struct {
	Page  int32 `json:"page"`
	Limit int32 `json:"limit"`
}

type PlayerFetchHistoryMsgResp struct {
	TotalPage   int32         `json:"total_page"`
	CurrentPage int32         `json:"current_page"`
	List        []interface{} `json:"list"`
}

type PlayerSendMsgReq struct {
	Content string `json:"content"`
	Pic     string `json:"pic"`
}

type CsFetchPlayerQueueReq struct {
	Limit int32 `json:"limit"`
}

type CsFetchPlayerQueueResp struct {
	List []interface{} `json:"list"`
}

type CsConnectPlayerReq struct {
	GameId   string `json:"game_id"`
	PlayerId string `json:"player_id"`
}

type CsFetchHistoryChatReq struct {
	Page  int32 `json:"page"`
	Limit int32 `json:"limit"`
}

type CsFetchHistoryChatResp struct {
	TotalPage   int32         `json:"total_page"`
	CurrentPage int32         `json:"current_page"`
	List        []interface{} `json:"list"`
}

type CsFetchHistoryMsgReq struct {
	GameId   string `json:"game_id"`
	PlayerId string `json:"player_id"`
	Page     int32  `json:"page"`
	Limit    int32  `json:"limit"`
}

type CsFetchHistoryMsgResp struct {
	TotalPage   int32         `json:"total_page"`
	CurrentPage int32         `json:"current_page"`
	List        []interface{} `json:"list"`
}

type CsFetchMsgReq struct {
	GameId   string `json:"game_id"`
	PlayerId string `json:"player_id"`
}

type CsSendMsgReq struct {
	GameId   string `json:"game_id"`
	PlayerId string `json:"player_id"`
	Content  string `json:"content"`
	Pic      string `json:"pic"`
}
