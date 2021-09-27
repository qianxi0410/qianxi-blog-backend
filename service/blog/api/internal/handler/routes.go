// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"qianxi-blog/service/blog/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/post/count",
				Handler: CountHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/post/count/:tag",
				Handler: CountWithTagHandler(serverCtx),
			},
		},
	)
}