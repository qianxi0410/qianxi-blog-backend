package logic

import (
	"context"
	"errors"
	"time"

	"qianxi-blog/common/key"
	"qianxi-blog/service/blog/api/internal/svc"
	"qianxi-blog/service/blog/api/internal/types"

	"github.com/go-redis/redis/v8"
	"github.com/tal-tech/go-zero/core/logx"
)

type CountWithTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCountWithTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) CountWithTagLogic {
	return CountWithTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CountWithTagLogic) CountWithTag(req types.CountWithTagReq) (*types.Reply, error) {
	if len(req.Tag) == 0 {
		return nil, errors.New("传入的标签为空")
	}

	count, err := l.svcCtx.Redis.Get(context.Background(), key.PostsCountWithTag(req.Tag)).Int64()

	if err != redis.Nil {
		return &types.Reply{
			Code: 666,
			Data: count,
		}, nil
	}

	count, err = l.svcCtx.PostModel.CountWtihTag(req.Tag)
	if err != nil {
		return nil, errors.New("查询带标签文章总数时出错" + err.Error())
	}

	l.svcCtx.Redis.Set(context.Background(), key.PostsCountWithTag(req.Tag), count, 1*time.Minute)

	return &types.Reply{
		Code: 666,
		Data: count,
	}, nil
}
