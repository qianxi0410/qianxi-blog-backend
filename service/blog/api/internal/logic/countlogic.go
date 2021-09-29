package logic

import (
	"context"
	"errors"
	"qianxi-blog/common/key"
	"time"

	"github.com/go-redis/redis/v8"

	"qianxi-blog/service/blog/api/internal/svc"
	"qianxi-blog/service/blog/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) CountLogic {
	return CountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CountLogic) Count() (*types.Reply, error) {
	var count int64

	count, err := l.svcCtx.Redis.Get(context.Background(), key.PostsCount()).Int64()

	if err != redis.Nil {
		return &types.Reply{
			Code: 666,
			Data: count,
		}, nil
	}

	count, err = l.svcCtx.PostModel.Count()
	if err != nil {
		return nil, errors.New("查询博客总数出错")
	}

	l.svcCtx.Redis.Set(context.Background(), key.PostsCount(), count, 1*time.Minute)

	return &types.Reply{
		Code: 666,
		Data: count,
	}, nil
}
