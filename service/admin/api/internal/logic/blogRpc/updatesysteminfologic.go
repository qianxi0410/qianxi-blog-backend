package blogRpc

import (
	"context"
	"qianxi-blog/service/blog/rpc/blog"

	"qianxi-blog/service/admin/api/internal/svc"
	"qianxi-blog/service/admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateSystemInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSystemInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateSystemInfoLogic {
	return UpdateSystemInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSystemInfoLogic) UpdateSystemInfo(req types.UpdateSystemReq) (*types.Reply, error) {
	_, err := l.svcCtx.BlogRpc.UpdateSystemInfo(l.ctx, &blog.UpdateSystemInfoReq{
		Key:   req.Key,
		Value: req.Value,
	})
	if err != nil {
		return nil, err
	}

	return &types.Reply{
		Code: 666,
	}, nil
}
