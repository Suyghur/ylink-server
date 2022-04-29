package handler

import (
	"net/http"
	"ylink/ext/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ylink/bff/cmdbff/api/internal/logic"
	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"
)

func playerFetchCsInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PlayerFetchCsInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewPlayerFetchCsInfoLogic(r.Context(), svcCtx)
		resp, err := l.PlayerFetchCsInfo(&req)
		result.HttpResult(r, w, resp, err)
	}
}
