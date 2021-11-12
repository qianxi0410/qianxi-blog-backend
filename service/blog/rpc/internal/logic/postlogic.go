package logic

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"qianxi-blog/common/key"
	"qianxi-blog/service/blog/model/wrapper"

	"github.com/go-redis/redis/v8"

	"qianxi-blog/service/blog/rpc/blog"
	"qianxi-blog/service/blog/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type PostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostLogic {
	return &PostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostLogic) Post(in *blog.PostReq) (*blog.PostReply, error) {
	var post = &wrapper.PostWrapper{}
	id := in.Id

	bytes, err := l.svcCtx.Redis.Get(context.Background(), key.Post(id)).Bytes()
	if err != redis.Nil {
		err := json.Unmarshal(bytes, &post)
		if err != nil {
			return nil, errors.New("查询文章时出错: " + err.Error())
		}
		return &blog.PostReply{Post: &blog.Post{
			Id:          post.Post.Id,
			CreatedAt:   post.Post.CreatedAt.Unix(),
			UpdatedAt:   post.Post.UpdatedAt.Unix(),
			Title:       post.Post.Title,
			Description: post.Post.Description.String,
			Pre:         post.Post.Pre,
			Next:        post.Post.Next,
			Url:         post.Post.Url,
			Path:        post.Post.Path,
			Tags:        post.Post.Tags.String,
		}}, nil
	}

	post.Post, err = l.svcCtx.PostModel.FindOne(id)
	if err != nil {
		return nil, errors.New("查询文章时出错: " + err.Error())
	}

	var bs []byte
	//bs, err = ioutil.ReadFile(post.Post.Path)
	bs, err = ioutil.ReadFile("E:\\idea_workspace\\qianxi-blog-backend-distribution\\README.md")

	if err != nil {
		panic(err)
	}
	post.Post.Path = string(bs)

	return &blog.PostReply{
		Post: &blog.Post{
			Id:          post.Post.Id,
			CreatedAt:   post.Post.CreatedAt.Unix(),
			UpdatedAt:   post.Post.UpdatedAt.Unix(),
			Title:       post.Post.Title,
			Description: post.Post.Description.String,
			Pre:         post.Post.Pre,
			Next:        post.Post.Next,
			Url:         post.Post.Url,
			Path:        post.Post.Path,
			Tags:        post.Post.Tags.String,
		},
	}, nil
}
