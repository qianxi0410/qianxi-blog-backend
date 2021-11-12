package logic

import (
	"context"
	"qianxi-blog/common/key"
	"time"

	"qianxi-blog/service/blog/rpc/blog"
	"qianxi-blog/service/blog/rpc/internal/svc"

	"github.com/go-redis/redis/v8"
	"github.com/tal-tech/go-zero/core/logx"
)

type CommentCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentCountLogic {
	return &CommentCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentCountLogic) CommentCount(in *blog.CountReq) (*blog.CountReply, error) {
	var count int64
	count, err := l.svcCtx.Redis.Get(l.ctx, key.CommentCount()).Int64()

	if err == redis.Nil {
		count, err = l.svcCtx.CommentModel.Count()

		if err != nil {
			return nil, err
		}

		l.svcCtx.Redis.Set(l.ctx, key.CommentCount(), count, 10*time.Minute)
	}

	return &blog.CountReply{
		Count: count,
	}, nil
}
