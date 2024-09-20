// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	app "oa/internal/handler/app"
	user "oa/internal/handler/user"
	"oa/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/login/:pa",
				Handler: app.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/app"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: user.RegHandler(serverCtx),
			},
			{
				Method:  http.MethodHead,
				Path:    "/:id",
				Handler: user.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/",
					Handler: user.ListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/:id",
					Handler: user.GetHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/user"),
	)
}
