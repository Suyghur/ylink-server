package cmd

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ylink/bff/apibff/api/internal/logic/player/cmd"
	"ylink/bff/apibff/api/internal/svc"
)

func PlayerDisconnectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := cmd.NewPlayerDisconnectLogic(r.Context(), svcCtx)
		resp, err := l.PlayerDisconnect()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
