package postApi

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"qianxi-blog/common/key"
	"qianxi-blog/service/blog/model/wrapper"
	"sync"
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
	var post = &wrapper.PostWrapper{}
	if req.Id < 1 {
		return nil, errors.New("文章id不合法")
	}

	bytes, err := l.svcCtx.Redis.Get(context.Background(), key.Post(req.Id)).Bytes()
	if err != redis.Nil {
		err := json.Unmarshal(bytes, &post)
		if err != nil {
			return nil, errors.New("查询文章时出错: " + err.Error())
		}
		return &types.Reply{Code: 666, Data: post}, nil
	}

	post.Post, err = l.svcCtx.PostModel.FindOne(req.Id)
	if err != nil {
		return nil, errors.New("查询文章时出错: " + err.Error())
	}

	var wg sync.WaitGroup
	wg.Add(1)
	var bs []byte
	go func() {
		defer wg.Done()
		// TODO:
		//bs, err = ioutil.ReadFile(post.Post.Path)
		bs, err = ioutil.ReadFile("E:\\idea_workspace\\qianxi-blog-backend-distribution\\README.md")
		if err != nil {
			panic(err)
		}
		post.Post.Path = string(bs)
	}()

	if post.Post.Pre != -1 {
		post.PreTitle, err = l.svcCtx.PostModel.Title(post.Post.Pre)
		if err != nil {
			return nil, errors.New("查询文章时出错: " + err.Error())
		}
	}

	if post.Post.Next != -1 {
		post.NextTitle, err = l.svcCtx.PostModel.Title(post.Post.Next)
		if err != nil {
			return nil, errors.New("查询文章时出错: " + err.Error())
		}
	}

	post.Comments, err = l.svcCtx.CommentModel.CommentsWithPostId(post.Post.Id)
	if err != nil {
		return nil, errors.New("查询文章评论时出错: " + err.Error())
	}

	wg.Wait()
	marshal, err := json.Marshal(post)
	if err != nil {
		return nil, errors.New("查询文章时出错: " + err.Error())
	}

	l.svcCtx.Redis.Set(context.Background(), key.Post(req.Id), marshal, 10*time.Minute)

	return &types.Reply{
		Code: 666,
		Data: post,
	}, nil
}
