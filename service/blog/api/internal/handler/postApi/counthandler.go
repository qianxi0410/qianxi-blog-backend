package postApi

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"qianxi-blog/service/blog/api/internal/logic/postApi"
	"qianxi-blog/service/blog/api/internal/svc"
)

func CountHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := postApi.NewCountLogic(r.Context(), ctx)
		resp, err := l.Count()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
