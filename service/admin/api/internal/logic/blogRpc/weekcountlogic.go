package blogRpc

import (
	"context"
	"qianxi-blog/service/blog/rpc/blogclient"

	"qianxi-blog/service/admin/api/internal/svc"
	"qianxi-blog/service/admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type WeekCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWeekCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) WeekCountLogic {
	return WeekCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WeekCountLogic) WeekCount() (*types.Reply, error) {
	count, err := l.svcCtx.BlogRpc.WeekCount(l.ctx, &blogclient.CountReq{})
	if err != nil {
		return nil, err
	}

	return &types.Reply{
		Code: 666,
		Data: count.Count,
	}, nil
}
