package cmd

import (
	"net/http"
	"ylink/ext/result"

	"ylink/bff/apibff/internal/logic/player/cmd"
	"ylink/bff/apibff/internal/svc"
)

func PlayerFetchMsgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := cmd.NewPlayerFetchMsgLogic(r.Context(), svcCtx)
		resp, err := l.PlayerFetchMsg()
		result.HttpResult(r, w, resp, err)
	}
}
