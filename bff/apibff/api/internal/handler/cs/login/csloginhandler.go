package login

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ylink/bff/apibff/api/internal/logic/cs/login"
	"ylink/bff/apibff/api/internal/svc"
	"ylink/bff/apibff/api/internal/types"
)

func CsLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CsLoginInfo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, &types.CommResp{
				Code: -1,
				Msg:  err.Error(),
				Data: map[string]interface{}{},
			})
			return
		}

		l := login.NewCsLoginLogic(r.Context(), svcCtx)
		resp, err := l.CsLogin(&req)
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
