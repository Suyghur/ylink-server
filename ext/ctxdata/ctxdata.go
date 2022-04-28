//@File     ctxdata.go
//@Time     2022/04/27
//@Author   #Suyghur,

package ctxdata

import (
	"context"
	"ylink/ext/jwtkey"
)

func GetPlayerIdFromCtx(ctx context.Context) string {
	playerId, _ := ctx.Value(jwtkey.PlayerId).(string)
	return playerId
}

func GetGameIdFromCtx(ctx context.Context) string {
	gameId, _ := ctx.Value(jwtkey.GameId).(string)
	return gameId
}

func GetCsIdFromCtx(ctx context.Context) string {
	csId, _ := ctx.Value(jwtkey.CsId).(string)
	return csId
}
