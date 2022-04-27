package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ylink/bff/authbff/internal/logic"
	"ylink/bff/authbff/internal/svc"
	"ylink/bff/authbff/internal/types"
)

func checkAuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckAuthReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCheckAuthLogic(r.Context(), svcCtx)
		resp, err := l.CheckAuth(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}