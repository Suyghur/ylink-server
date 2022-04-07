package cmd

import (
	"net/http"
	"ylink/bff/apibff/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ylink/bff/apibff/api/internal/logic/cs/cmd"
	"ylink/bff/apibff/api/internal/svc"
)

func CsLogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := cmd.NewCsLogoutLogic(r.Context(), svcCtx)
		resp, err := l.CsLogout()
		if err != nil {
			httpx.OkJson(w, &types.CommResp{
				Code: -1,
				Msg:  err.Error(),
				Data: map[string]interface{}{},
			})
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
