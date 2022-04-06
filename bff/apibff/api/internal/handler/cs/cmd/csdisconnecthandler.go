package cmd

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ylink/bff/apibff/api/internal/logic/cs/cmd"
	"ylink/bff/apibff/api/internal/svc"
	"ylink/bff/apibff/api/internal/types"
)

func CsDisconnectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CsDisconnectReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := cmd.NewCsDisconnectLogic(r.Context(), svcCtx)
		resp, err := l.CsDisconnect(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
