package blogRpc

import (
	"context"
	"encoding/json"
	"fmt"
	"qianxi-blog/common/key"
	"qianxi-blog/service/blog/rpc/blogclient"
	"time"

	"github.com/go-redis/redis/v8"

	"qianxi-blog/service/admin/api/internal/svc"
	"qianxi-blog/service/admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CountInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCountInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) CountInfoLogic {
	return CountInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CountInfoLogic) CountInfo() (*types.Reply, error) {
	var res = make([]int64, 4)

	jsonBytes, err := l.svcCtx.Redis.Get(l.ctx, key.CountInfo()).Bytes()
	fmt.Println(jsonBytes)
	fmt.Println(err)
	if err != redis.Nil {
		err = json.Unmarshal(jsonBytes, &res)
		if err != nil {
			return nil, err
		}

		return &types.Reply{Code: 666, Data: res}, nil
	}

	postCount, err := l.svcCtx.BlogRpc.PostCount(l.ctx, &blogclient.CountReq{})
	if err != nil {
		return nil, err
	}
	res[0] = postCount.Count

	commentCount, err := l.svcCtx.BlogRpc.CommentCount(l.ctx, &blogclient.CountReq{})
	if err != nil {
		return nil, err
	}

	res[1] = commentCount.Count

	visitedCount, err := l.svcCtx.BlogRpc.VisitedCount(l.ctx, &blogclient.CountReq{})
	if err != nil {
		return nil, err
	}

	res[2] = visitedCount.Count

	peopleCount, err := l.svcCtx.BlogRpc.PeopleCount(l.ctx, &blogclient.CountReq{})
	if err != nil {
		return nil, err
	}

	res[3] = peopleCount.Count

	marshal, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	l.svcCtx.Redis.Set(l.ctx, key.CountInfo(), marshal, 10*time.Minute)

	return &types.Reply{
		Code: 666,
		Data: res,
	}, nil
}
