package handler

import (
	"net/http"
	"ylink/ext/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ylink/bff/cmdbff/api/internal/logic"
	"ylink/bff/cmdbff/api/internal/svc"
	"ylink/bff/cmdbff/api/internal/types"
)

func csSendMsgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CsSendMsgReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewCsSendMsgLogic(r.Context(), svcCtx)
		err := l.CsSendMsg(&req)
		result.HttpResult(r, w, nil, err)
	}
}
