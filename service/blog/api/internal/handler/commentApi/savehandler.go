package commentApi

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"qianxi-blog/service/blog/api/internal/logic/commentApi"
	"qianxi-blog/service/blog/api/internal/svc"
)

func SaveHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := commentApi.NewSaveLogic(r.Context(), ctx)
		resp, err := l.Save()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
