package player

import (
	"net/http"
	"ylink/ext/result"

	"ylink/bff/cmdbff/api/internal/logic/player"
	"ylink/bff/cmdbff/api/internal/svc"
)

func DisconnectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := player.NewDisconnectLogic(r.Context(), svcCtx)
		err := l.Disconnect()
		result.HttpResult(r, w, nil, err)
	}
}
