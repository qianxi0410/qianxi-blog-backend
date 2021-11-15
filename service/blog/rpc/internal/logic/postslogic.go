package logic

import (
	"context"
	"encoding/json"
	"qianxi-blog/common/key"
	"qianxi-blog/service/blog/model"
	"qianxi-blog/service/blog/rpc/blog"
	"qianxi-blog/service/blog/rpc/internal/svc"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/tal-tech/go-zero/core/logx"
)

type PostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostsLogic {
	return &PostsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostsLogic) Posts(in *blog.PageReq) (*blog.PageReply, error) {
	page, size := in.Page, in.Size

	var posts []model.Posts
	var reply []*blog.Post

	// all article
	if size == -1 {
		all, err := l.svcCtx.PostModel.PostsAll()
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(all); i++ {
			reply = append(reply, &blog.Post{
				Id:          all[i].Id,
				CreatedAt:   all[i].CreatedAt.Unix(),
				UpdatedAt:   all[i].UpdatedAt.Unix(),
				Title:       all[i].Title,
				Description: all[i].Description.String,
				Pre:         all[i].Pre,
				Next:        all[i].Next,
				Url:         all[i].Url,
				Path:        all[i].Path,
				Tags:        all[i].Tags.String,
				Blur:        all[i].Blur,
			})
		}

		return &blog.PageReply{
			Posts: reply,
		}, nil
	}

	jsonBytes, err := l.svcCtx.Redis.Get(context.Background(), key.Posts(page, size)).Bytes()

	if err != redis.Nil {
		err := json.Unmarshal(jsonBytes, &posts)
		if err != nil {
			return nil, err
		}

	} else {
		offset := (page - 1) * size
		posts, err = l.svcCtx.PostModel.Posts(offset, size)

		if err != nil {
			return nil, err
		}

		marshal, err := json.Marshal(posts)

		if err != nil {
			return nil, err
		}

		l.svcCtx.Redis.Set(l.ctx, key.Posts(page, size), marshal, 10*time.Minute)
	}

	for i := 0; i < len(posts); i++ {
		reply = append(reply, &blog.Post{
			Id:          posts[i].Id,
			CreatedAt:   posts[i].CreatedAt.Unix(),
			UpdatedAt:   posts[i].UpdatedAt.Unix(),
			Title:       posts[i].Title,
			Description: posts[i].Description.String,
			Pre:         posts[i].Pre,
			Next:        posts[i].Next,
			Url:         posts[i].Url,
			Path:        posts[i].Path,
			Tags:        posts[i].Tags.String,
			Blur:        posts[i].Blur,
		})
	}

	return &blog.PageReply{
		Posts: reply,
	}, nil
}
