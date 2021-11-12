package postApi

import (
	"context"
	"encoding/json"
	"errors"
	"qianxi-blog/common/key"
	"qianxi-blog/service/blog/api/internal/svc"
	"qianxi-blog/service/blog/api/internal/types"
	"qianxi-blog/service/blog/model"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/tal-tech/go-zero/core/logx"
)

type PostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) PostsLogic {
	return PostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostsLogic) Posts(req types.PageReq) (*types.Reply, error) {
	var posts []model.Posts

	if req.Page <= 0 || req.Size <= 0 {
		return nil, errors.New("分页查询的参数不能为负")
	}

	jsonBytes, err := l.svcCtx.Redis.Get(context.Background(), key.Posts(req.Page, req.Size)).Bytes()

	if err != redis.Nil {
		err := json.Unmarshal(jsonBytes, &posts)
		if err != nil {
			return nil, errors.New("分页查询文章时出错: " + err.Error())
		}
		return &types.Reply{
			Code: 666,
			Data: posts,
		}, nil
	}

	offset := (req.Page - 1) * req.Size

	posts, err = l.svcCtx.PostModel.Posts(offset, req.Size)
	if err != nil {
		return nil, errors.New("分页查询文章时出错: " + err.Error())
	}

	marshal, err := json.Marshal(posts)

	if err != nil {
		return nil, errors.New("序列化文章时出错: " + err.Error())
	}

	l.svcCtx.Redis.Set(context.Background(), key.Posts(req.Page, req.Size), marshal, 10*time.Minute)

	return &types.Reply{
		Code: 666,
		Data: posts,
	}, nil
}
