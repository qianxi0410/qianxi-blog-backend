package systemApi

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"qianxi-blog/service/blog/api/internal/logic/systemApi"
	"qianxi-blog/service/blog/api/internal/svc"
)

func SystemInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := systemApi.NewSystemInfoLogic(r.Context(), ctx)
		resp, err := l.SystemInfo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
