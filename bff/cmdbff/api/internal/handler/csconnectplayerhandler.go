package handler

import (
	"net/http"
	"ylink/ext/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ylink/bff/cmdbff/api/internal/logic"
	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"
)

func csConnectPlayerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CsConnectPlayerReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewCsConnectPlayerLogic(r.Context(), svcCtx)
		err := l.CsConnectPlayer(&req)
		result.HttpResult(r, w, nil, err)
	}
}
