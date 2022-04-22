package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ylink/bff/authbff/internal/logic"
	"ylink/bff/authbff/internal/svc"
	"ylink/bff/authbff/internal/types"
)

func csAuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CsAuthReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, &types.AuthResp{
				Code: 1,
				Msg:  err.Error(),
				Data: map[string]interface{}{},
			})
			return
		}

		l := logic.NewCsAuthLogic(r.Context(), svcCtx)
		resp, err := l.CsAuth(&req)
		if err != nil {
			httpx.OkJson(w, &types.AuthResp{
				Code: 1,
				Msg:  err.Error(),
				Data: map[string]interface{}{},
			})
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
