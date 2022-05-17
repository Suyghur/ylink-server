package handler

import (
	"net/http"
	"ylink/bff/authbff/api/internal/logic"
	"ylink/bff/authbff/api/internal/svc"
	"ylink/bff/authbff/api/internal/types"
	"ylink/comm/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func playerLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PlayerAuthReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewPlayerLoginLogic(r.Context(), svcCtx)
		resp, err := l.PlayerLogin(&req)
		result.HttpResult(r, w, resp, err)
	}
}
