package logic

import (
	"context"
	"encoding/json"
	"qianxi-blog/common/key"
	"time"

	"github.com/go-redis/redis/v8"

	"qianxi-blog/service/blog/rpc/blog"
	"qianxi-blog/service/blog/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type WeekCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWeekCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WeekCountLogic {
	return &WeekCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WeekCountLogic) WeekCount(in *blog.CountReq) (*blog.WeekCountReply, error) {
	result, err := l.svcCtx.Redis.Get(l.ctx, key.WeekCount()).Result()
	if err != redis.Nil {
		var r []int64
		err := json.Unmarshal([]byte(result), &r)
		if err != nil {
			return nil, err
		}
		return &blog.WeekCountReply{Count: r}, nil
	}

	weekCounts, err := l.svcCtx.VisitModel.WeekCount()
	if err != nil {
		return nil, err
	}

	counts := make([]int64, 7)

	for i := 6; i >= 0; i-- {
		day := time.Now().AddDate(0, 0, -i).Day()
		for j := 0; j < len(weekCounts); j++ {
			if weekCounts[j].Day == int64(day) {
				counts[6-i] = weekCounts[j].Count
				break
			}
		}
	}

	marshal, err := json.Marshal(counts)
	if err != nil {
		return nil, err
	}
	l.svcCtx.Redis.Set(l.ctx, key.WeekCount(), marshal, 10*time.Minute)

	return &blog.WeekCountReply{
		Count: counts,
	}, nil
}
