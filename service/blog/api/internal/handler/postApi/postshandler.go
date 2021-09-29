package postApi

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"qianxi-blog/service/blog/api/internal/logic/postApi"
	"qianxi-blog/service/blog/api/internal/svc"
	"qianxi-blog/service/blog/api/internal/types"
)

func PostsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := postApi.NewPostsLogic(r.Context(), ctx)
		resp, err := l.Posts(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
