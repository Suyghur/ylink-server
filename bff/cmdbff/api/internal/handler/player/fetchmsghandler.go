package player

import (
	"net/http"
	"ylink/ext/result"

	"ylink/bff/cmdbff/api/internal/logic/player"
	"ylink/bff/cmdbff/api/internal/svc"
)

func FetchMsgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := player.NewFetchMsgLogic(r.Context(), svcCtx)
		resp, err := l.FetchMsg()
		result.HttpResult(r, w, resp, err)
	}
}
