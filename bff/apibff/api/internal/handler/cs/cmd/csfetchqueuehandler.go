package cmd

import (
	"net/http"
	"ylink/bff/apibff/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ylink/bff/apibff/api/internal/logic/cs/cmd"
	"ylink/bff/apibff/api/internal/svc"
)

func CsFetchQueueHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := cmd.NewCsFetchQueueLogic(r.Context(), svcCtx)
		resp, err := l.CsFetchQueue()
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
