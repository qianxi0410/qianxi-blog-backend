package logic

import (
	"context"
	"database/sql"
	"qianxi-blog/common/key"
	"qianxi-blog/service/blog/rpc/internal/utils"
	"sync"
	"time"

	"qianxi-blog/service/blog/model"
	"qianxi-blog/service/blog/rpc/blog"
	"qianxi-blog/service/blog/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateWithContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateWithContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateWithContentLogic {
	return &UpdateWithContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateWithContentLogic) UpdateWithContent(in *blog.UpdateReq) (*blog.UpdateReply, error) {
	var wg sync.WaitGroup
	wg.Add(2)

	one, err := l.svcCtx.PostModel.FindOne(in.Post.Id)
	if err != nil {
		return nil, err
	}

	path := one.Path
	content := in.Post.Path

	go func() {
		defer wg.Done()

		err = utils.WriteFile(path, []byte(content), 0766)
		if err != nil {
			panic(err)
		}
	}()

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
		Path: path,
		Tags: sql.NullString{
			Valid:  true,
			String: in.Post.Tags,
		},
	}

	go func() {
		defer wg.Done()
		l.svcCtx.Redis.Expire(l.ctx, key.Post(post.Id), 0)
		result, err := l.svcCtx.Redis.Keys(l.ctx, "*_*_page_size").Result()
		if err != nil {
			panic(err)
		}
		for i := 0; i < len(result); i++ {
			l.svcCtx.Redis.Expire(l.ctx, result[i], 0)
		}
	}()

	err = l.svcCtx.PostModel.Update(post)
	if err != nil {
		wg.Wait()
		return nil, err
	}

	wg.Wait()
	return &blog.UpdateReply{}, nil
}
