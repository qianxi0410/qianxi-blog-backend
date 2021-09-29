package logic

import (
	"context"
	"encoding/json"
	"errors"
	"qianxi-blog/common/key"
	"qianxi-blog/service/blog/model"
	"time"

	"github.com/go-redis/redis/v8"

	"qianxi-blog/service/blog/api/internal/svc"
	"qianxi-blog/service/blog/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type PostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) PostLogic {
	return PostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostLogic) Post(req types.PostReq) (*types.Reply, error) {
	var post *model.Posts
	if req.Id < 1 {
		return nil, errors.New("文章id不合法")
	}

	bytes, err := l.svcCtx.Redis.Get(context.Background(), key.Post(req.Id)).Bytes()
	if err != redis.Nil {
		err := json.Unmarshal(bytes, &post)
		if err != nil {
			return nil, errors.New("查询博客时出错: " + err.Error())
		}
		return &types.Reply{Code: 666, Data: post}, nil
	}

	post, err = l.svcCtx.PostModel.FindOne(req.Id)
	if err != nil {
		return nil, errors.New("查询博客时出错: " + err.Error())
	}

	marshal, _ := json.Marshal(post)

	l.svcCtx.Redis.Set(context.Background(), key.Post(req.Id), marshal, 10*time.Minute)

	return &types.Reply{
		Code: 666,
		Data: post,
	}, nil
}
