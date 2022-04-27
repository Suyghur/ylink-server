package cmd

import (
	"net/http"
	"ylink/ext/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ylink/bff/apibff/internal/logic/cs/cmd"
	"ylink/bff/apibff/internal/svc"
	"ylink/bff/apibff/internal/types"
)

func CsConnectPlayerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CsConnectPlayerReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
		}

		l := cmd.NewCsConnectPlayerLogic(r.Context(), svcCtx)
		err := l.CsConnectPlayer(&req)
		result.HttpResult(r, w, nil, err)
	}
}
