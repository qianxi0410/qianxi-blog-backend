package logic

import (
	"context"
	"database/sql"
	"qianxi-blog/common/key"
	"qianxi-blog/service/blog/model"
	"sync"
	"time"

	"qianxi-blog/service/blog/rpc/blog"
	"qianxi-blog/service/blog/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *blog.UpdateReq) (*blog.UpdateReply, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	post := model.Posts{
		Id:        in.Post.Id,
		CreatedAt: time.Unix(in.Post.CreatedAt, 0),
		UpdatedAt: time.Now(),
		Title:     in.Post.Title,
		Description: sql.NullString{
			Valid:  true,
			String: in.Post.Description,
		},
		Pre:  in.Post.Pre,
		Next: in.Post.Next,
		Url:  in.Post.Url,
		Path: in.Post.Path,
		Tags: sql.NullString{
			Valid:  true,
			String: in.Post.Tags,
		},
		Blur: in.Post.Blur,
	}

	go func() {
		defer wg.Done()
		l.svcCtx.Redis.Expire(l.ctx, key.Post(post.Id), 0)
		l.svcCtx.Redis.Expire(l.ctx, key.Posts(in.Page, in.Size), 0)
	}()

	err := l.svcCtx.PostModel.Update(post)
	if err != nil {
		wg.Wait()
		return nil, err
	}

	wg.Wait()
	return &blog.UpdateReply{}, nil
}
