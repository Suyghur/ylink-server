package handler

import (
	"net/http"
	"ylink/ext/result"

	"ylink/bff/cmdbff/api/internal/logic"
	"ylink/bff/cmdbff/api/internal/svc"
)

func playerDisconnectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewPlayerDisconnectLogic(r.Context(), svcCtx)
		err := l.PlayerDisconnect()
		result.HttpResult(r, w, nil, err)
	}
}
