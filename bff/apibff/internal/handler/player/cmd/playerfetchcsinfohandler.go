package cmd

import (
	"net/http"
	"ylink/ext/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ylink/bff/apibff/internal/logic/player/cmd"
	"ylink/bff/apibff/internal/svc"
	"ylink/bff/apibff/internal/types"
)

func PlayerFetchCsInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PlayerFetchCsInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
		}

		l := cmd.NewPlayerFetchCsInfoLogic(r.Context(), svcCtx)
		resp, err := l.PlayerFetchCsInfo(&req)
		result.HttpResult(r, w, resp, err)
	}
}
