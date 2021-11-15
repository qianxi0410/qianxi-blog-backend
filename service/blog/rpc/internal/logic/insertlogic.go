package logic

import (
	"context"
	"database/sql"
	"qianxi-blog/common/key"
	"qianxi-blog/service/blog/model"
	"qianxi-blog/service/blog/rpc/internal/utils"
	"strings"
	"time"

	"qianxi-blog/service/blog/rpc/blog"
	"qianxi-blog/service/blog/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type InsertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertLogic {
	return &InsertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertLogic) Insert(in *blog.InsertReq) (*blog.InsertReply, error) {
	path := ""
	// TODO: ./ => /blog
	if strings.Contains(in.Title, " ") {
		path = "./" + strings.Join(strings.Split(in.Title, " "), "-") + ".md"
	} else {
		path = "./" + in.Title + ".md"
	}

	id, err := l.svcCtx.PostModel.MaxId()
	if err != nil {
		return nil, err
	}

	one, err := l.svcCtx.PostModel.FindOne(id)
	if err != nil {
		return nil, err
	}
	one.Pre = id + 1
	err = l.svcCtx.PostModel.Update(*one)
	if err != nil {
		return nil, err
	}

	post := model.Posts{
		Id:        id + 1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Title:     in.Title,
		Description: sql.NullString{
			Valid:  true,
			String: in.Description,
		},
		Pre:  -1,
		Next: id,
		Url:  in.Url,
		Path: path,
		Tags: sql.NullString{
			Valid:  true,
			String: in.Tags,
		},
		Blur: in.Blur,
	}

	err = utils.WriteFile(path, []byte(in.Path), 0766)
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.PostModel.Insert(post)
	if err != nil {
		return nil, err
	}

	l.svcCtx.Redis.Expire(l.ctx, key.CountInfo(), 0)
	l.svcCtx.Redis.Expire(l.ctx, key.PostsCount(), 0)
	result, err := l.svcCtx.Redis.Keys(l.ctx, "*_*_page_size").Result()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(result); i++ {
		l.svcCtx.Redis.Expire(l.ctx, result[i], 0)
	}

	return &blog.InsertReply{}, nil
}
