package handler

import (
	"net/http"
	"ylink/ext/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ylink/bff/cmdbff/api/internal/logic"
	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"
)

func playerSendMsgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PlayerSendMsgReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewPlayerSendMsgLogic(r.Context(), svcCtx)
		err := l.PlayerSendMsg(&req)
		result.HttpResult(r, w, nil, err)
	}
}
