package handler

import (
	"net/http"

	"call_center/db/api/internal/logic"
	"call_center/db/api/internal/svc"
	"call_center/db/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func apiTestHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewApiTestLogic(r.Context(), ctx)
		resp, err := l.ApiTest(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
