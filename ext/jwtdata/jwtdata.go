//@File     jwtdata.go
//@Time     2022/04/24
//@Author   #Suyghur,

package jwtdata

import "context"

const (
	JwtKeyPlayerId = "jwt_player_id"
	JwtKeyGameId   = "jwt_game_id"
	JwtKeyCsId     = "jwt_cs_id"
)

func GetPlayerIdFromJwt(ctx context.Context) string {
	playerId, _ := ctx.Value(JwtKeyPlayerId).(string)
	return playerId
}

func GetGameIdFromJwt(ctx context.Context) string {
	gameId, _ := ctx.Value(JwtKeyGameId).(string)
	return gameId
}

func GetCsIdFromJwt(ctx context.Context) string {
	csId, _ := ctx.Value(JwtKeyCsId).(string)
	return csId
}
