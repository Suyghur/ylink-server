// Code generated by goctl. DO NOT EDIT.
package types

type PlayerAuthReq struct {
	GameId   string `json:"game_id"`
	PlayerId string `json:"player_id"`
	Type     int32  `json:"type"`
}

type CsAuthReq struct {
	CsId string `json:"cs_id"`
}

type AuthResp struct {
	AccessToken string `json:"access_token"`
}
