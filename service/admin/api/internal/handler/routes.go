// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	userApi "qianxi-blog/service/admin/api/internal/handler/userApi"
	"qianxi-blog/service/admin/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/admin/login",
				Handler: userApi.LoginHandler(serverCtx),
			},
		},
	)
}
