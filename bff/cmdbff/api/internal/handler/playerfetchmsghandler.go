package handler

import (
	"net/http"
	"ylink/ext/result"

	"ylink/bff/cmdbff/api/internal/logic"
	"ylink/bff/cmdbff/api/internal/svc"
)

func playerFetchMsgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewPlayerFetchMsgLogic(r.Context(), svcCtx)
		resp, err := l.PlayerFetchMsg()
		result.HttpResult(r, w, resp, err)
	}
}
