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

type PostCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostCountLogic {
	return &PostCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostCountLogic) PostCount(in *blog.CountReq) (*blog.CountReply, error) {
	var count int64
	count, err := l.svcCtx.Redis.Get(l.ctx, key.PostsCount()).Int64()

	if err == redis.Nil {
		count, err = l.svcCtx.PostModel.Count()

		if err != nil {
			return nil, err
		}

		l.svcCtx.Redis.Set(l.ctx, key.PostsCount(), count, 10*time.Minute)
	}

	return &blog.CountReply{
		Count: count,
	}, nil
}
