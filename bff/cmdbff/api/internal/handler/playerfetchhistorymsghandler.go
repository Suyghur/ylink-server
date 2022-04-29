package handler

import (
	"net/http"
	"ylink/ext/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ylink/bff/cmdbff/api/internal/logic"
	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"
)

func playerFetchHistoryMsgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PlayerFetchHistoryMsgReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewPlayerFetchHistoryMsgLogic(r.Context(), svcCtx)
		resp, err := l.PlayerFetchHistoryMsg(&req)
		result.HttpResult(r, w, resp, err)
	}
}
