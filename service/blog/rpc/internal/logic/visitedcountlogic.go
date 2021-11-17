package logic

import (
	"context"
	"qianxi-blog/common/key"
	"time"

	"github.com/go-redis/redis/v8"

	"qianxi-blog/service/blog/rpc/blog"
	"qianxi-blog/service/blog/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type VisitedCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVisitedCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VisitedCountLogic {
	return &VisitedCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VisitedCountLogic) VisitedCount(in *blog.CountReq) (*blog.CountReply, error) {
	i, err := l.svcCtx.Redis.Get(l.ctx, key.VisitedCount()).Int64()
	if err != redis.Nil {
		return &blog.CountReply{Count: i}, nil
	}

	count, err := l.svcCtx.VisitModel.Count()
	if err != nil {
		return nil, err
	}

	l.svcCtx.Redis.Set(l.ctx, key.VisitedCount(), count, 10*time.Minute)

	return &blog.CountReply{
		Count: count,
	}, nil
}
