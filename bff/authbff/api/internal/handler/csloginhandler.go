package handler

import (
	"net/http"
	"ylink/bff/authbff/api/internal/logic"
	"ylink/bff/authbff/api/internal/svc"
	"ylink/bff/authbff/api/internal/types"
	"ylink/ext/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func csLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CsAuthReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewCsLoginLogic(r.Context(), svcCtx)
		resp, err := l.CsLogin(&req)
		result.HttpResult(r, w, resp, err)
	}
}
