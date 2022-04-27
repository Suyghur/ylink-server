package cmd

import (
	"net/http"
	"ylink/ext/result"

	"ylink/bff/apibff/internal/logic/player/cmd"
	"ylink/bff/apibff/internal/svc"
)

func PlayerDisconnectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := cmd.NewPlayerDisconnectLogic(r.Context(), svcCtx)
		err := l.PlayerDisconnect()
		result.HttpResult(r, w, nil, err)
	}
}
