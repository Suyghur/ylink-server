//@File     ctxdata.go
//@Time     2022/04/27
//@Author   #Suyghur,

package ctxdata

import (
	"context"
	"go.opentelemetry.io/otel/trace"
	"ylink/comm/jwtkey"
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

func GetTraceIdFromCtx(ctx context.Context) string {
	spanCtx := trace.SpanContextFromContext(ctx)
	if spanCtx.HasTraceID() {
		return spanCtx.TraceID().String()
	}
	return ""
}
