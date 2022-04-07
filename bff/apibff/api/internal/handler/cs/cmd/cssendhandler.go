package cmd

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ylink/bff/apibff/api/internal/logic/cs/cmd"
	"ylink/bff/apibff/api/internal/svc"
	"ylink/bff/apibff/api/internal/types"
)

func CsSendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatMsgReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, &types.CommResp{
				Code: -1,
				Msg:  err.Error(),
				Data: map[string]interface{}{},
			})
			return
		}

		l := cmd.NewCsSendLogic(r.Context(), svcCtx)
		resp, err := l.CsSend(&req)
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
