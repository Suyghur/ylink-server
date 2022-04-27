//@File     ctxdata.go
//@Time     2022/04/27
//@Author   #Suyghur,

package ctxdata

import (
	"context"
	"ylink/ext/jwtdata"
)

func GetPlayerIdFromCtx(ctx context.Context) string {
	playerId, _ := ctx.Value(jwtdata.JwtKeyPlayerId).(string)
	return playerId
}

func GetGameIdFromCtx(ctx context.Context) string {
	gameId, _ := ctx.Value(jwtdata.JwtKeyGameId).(string)
	return gameId
}

func GetCsIdFromJwt(ctx context.Context) string {
	csId, _ := ctx.Value(jwtdata.JwtKeyCsId).(string)
	return csId
}
