package userApi

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"qianxi-blog/service/admin/api/internal/logic/userApi"
	"qianxi-blog/service/admin/api/internal/svc"
)

func LoginValidHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := userApi.NewLoginValidLogic(r.Context(), ctx)
		resp, err := l.LoginValid()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
