package blogRpc

import (
	"context"
	"qianxi-blog/service/blog/rpc/blogclient"

	"qianxi-blog/service/admin/api/internal/svc"
	"qianxi-blog/service/admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AllSystemInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllSystemInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) AllSystemInfoLogic {
	return AllSystemInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllSystemInfoLogic) AllSystemInfo() (*types.Reply, error) {
	all, err := l.svcCtx.BlogRpc.SystemInfoAll(l.ctx, &blogclient.SystemInfoAllReq{})

	if err != nil {
		return nil, err
	}

	return &types.Reply{
		Code: 666,
		Data: all.Kv,
	}, nil
}
