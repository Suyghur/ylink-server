// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"ylink/bff/cmdbff/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/player/fetch-cs-info",
				Handler: playerFetchCsInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/player/fetch-history-msg",
				Handler: playerFetchHistoryMsgHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/player/fetch-msg",
				Handler: playerFetchMsgHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/player/send-msg",
				Handler: playerSendMsgHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/player/disconnect",
				Handler: playerDisconnectHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/v1/cmd"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/cs/fetch-player-queue",
				Handler: csFetchPlayerQueueHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/cs/connect-player",
				Handler: csConnectPlayerHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/cs/fetch-history-list",
				Handler: csFetchHistoryListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/cs/fetch-history-msg",
				Handler: csFetchHistoryMsgHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/cs/fetch-msg",
				Handler: csFetchMsgHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/cs/send-msg",
				Handler: csSendMsgHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/v1/cmd"),
	)
}
