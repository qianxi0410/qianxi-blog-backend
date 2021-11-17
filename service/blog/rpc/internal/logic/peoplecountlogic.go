package logic

import (
	"context"
	"time"

	"qianxi-blog/common/key"
	"qianxi-blog/service/blog/rpc/blog"
	"qianxi-blog/service/blog/rpc/internal/svc"

	"github.com/go-redis/redis/v8"
	"github.com/tal-tech/go-zero/core/logx"
)

type PeopleCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPeopleCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PeopleCountLogic {
	return &PeopleCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PeopleCountLogic) PeopleCount(in *blog.CountReq) (*blog.CountReply, error) {
	i, err := l.svcCtx.Redis.Get(l.ctx, key.PeopleCount()).Int64()
	if err != redis.Nil {
		return &blog.CountReply{Count: i}, nil
	}

	count, err := l.svcCtx.VisitModel.PeopleCount()
	if err != nil {
		return nil, err
	}

	l.svcCtx.Redis.Set(l.ctx, key.PeopleCount(), count, 10*time.Minute)

	return &blog.CountReply{
		Count: count,
	}, nil
}
