package blogRpc

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"qianxi-blog/service/admin/api/internal/logic/blogRpc"
	"qianxi-blog/service/admin/api/internal/svc"
	"qianxi-blog/service/admin/api/internal/types"
)

func UpdateSystemInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateSystemReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := blogRpc.NewUpdateSystemInfoLogic(r.Context(), ctx)
		resp, err := l.UpdateSystemInfo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
