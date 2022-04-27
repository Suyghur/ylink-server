package cmd

import (
	"net/http"
	"ylink/ext/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ylink/bff/apibff/internal/logic/player/cmd"
	"ylink/bff/apibff/internal/svc"
	"ylink/bff/apibff/internal/types"
)

func PlayerSendMsgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PlayerSendMsgReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
		}

		l := cmd.NewPlayerSendMsgLogic(r.Context(), svcCtx)
		err := l.PlayerSendMsg(&req)
		result.HttpResult(r, w, nil, err)
	}
}
