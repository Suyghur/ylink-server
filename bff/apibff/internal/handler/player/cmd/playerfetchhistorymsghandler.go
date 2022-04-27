package cmd

import (
	"net/http"
	"ylink/ext/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ylink/bff/apibff/internal/logic/player/cmd"
	"ylink/bff/apibff/internal/svc"
	"ylink/bff/apibff/internal/types"
)

func PlayerFetchHistoryMsgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PlayerFetchHistoryMsgReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
		}

		l := cmd.NewPlayerFetchHistoryMsgLogic(r.Context(), svcCtx)
		resp, err := l.PlayerFetchHistoryMsg(&req)
		result.HttpResult(r, w, resp, err)
	}
}
