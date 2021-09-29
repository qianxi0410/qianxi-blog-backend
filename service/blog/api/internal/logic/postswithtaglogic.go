package logic

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"qianxi-blog/common/key"
	"qianxi-blog/service/blog/api/internal/svc"
	"qianxi-blog/service/blog/api/internal/types"
	"qianxi-blog/service/blog/model"

	"github.com/go-redis/redis/v8"
	"github.com/tal-tech/go-zero/core/logx"
)

type PostsWithTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostsWithTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) PostsWithTagLogic {
	return PostsWithTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostsWithTagLogic) PostsWithTag(req types.PageWithTagReq) (*types.Reply, error) {
	var posts []model.Posts

	if req.Page <= 0 || req.Size <= 0 {
		return nil, errors.New("分页查询的参数不能为负")
	}

	if len(req.Tag) == 0 {
		return nil, errors.New("传入的标签不能为空")
	}

	bytes, err := l.svcCtx.Redis.Get(context.Background(), key.PostsWithTag(req.Page, req.Size, req.Tag)).Bytes()
	if err != redis.Nil {
		err := json.Unmarshal(bytes, &posts)
		if err != nil {
			return nil, errors.New("分页查询文章时出错: " + err.Error())
		}
		return &types.Reply{
			Code: 666,
			Data: posts,
		}, nil
	}

	offset := (req.Page - 1) * req.Size

	posts, err = l.svcCtx.PostModel.PostsWithTag(offset, req.Size, req.Tag)

	if err != nil {
		return nil, errors.New("分页查询文章时出错: " + err.Error())
	}

	marshal, err := json.Marshal(posts)
	if err != nil {
		return nil, errors.New("分页查询文章时出错: " + err.Error())
	}

	l.svcCtx.Redis.Set(context.Background(), key.PostsWithTag(req.Page, req.Size, req.Tag), marshal, 10*time.Minute)

	return &types.Reply{
		Code: 666,
		Data: posts,
	}, nil
}
