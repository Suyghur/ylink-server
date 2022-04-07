package cmd

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ylink/bff/apibff/api/internal/logic/player/cmd"
	"ylink/bff/apibff/api/internal/svc"
	"ylink/bff/apibff/api/internal/types"
)

func PlayerConnectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PlayerConnectReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, &types.CommResp{
				Code: -1,
				Msg:  err.Error(),
				Data: map[string]interface{}{},
			})
			return
		}

		l := cmd.NewPlayerConnectLogic(r.Context(), svcCtx)
		resp, err := l.PlayerConnect(&req)
		if err != nil {
			httpx.OkJson(w, &types.CommResp{
				Code: -1,
				Msg:  err.Error(),
				Data: map[string]interface{}{},
			})
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
