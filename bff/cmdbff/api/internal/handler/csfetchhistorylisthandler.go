package handler

import (
	"net/http"
	"ylink/comm/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ylink/bff/cmdbff/api/internal/logic"
	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"
)

func csFetchHistoryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CsFetchHistoryChatReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewCsFetchHistoryListLogic(r.Context(), svcCtx)
		resp, err := l.CsFetchHistoryList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
