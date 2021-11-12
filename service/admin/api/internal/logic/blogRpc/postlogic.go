package blogRpc

import (
	"context"
	"database/sql"
	"errors"
	"qianxi-blog/service/admin/model"
	"qianxi-blog/service/blog/rpc/blog"
	"time"

	"qianxi-blog/service/admin/api/internal/svc"
	"qianxi-blog/service/admin/api/internal/types"

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
	id := req.Id

	if id < 1 {
		return nil, errors.New("文章id不合法")
	}

	post, err := l.svcCtx.BlogRpc.Post(l.ctx, &blog.PostReq{Id: id})
	if err != nil {
		return nil, err
	}

	p := model.Posts{
		Id:        post.Post.Id,
		CreatedAt: time.Unix(post.Post.CreatedAt, 0),
		UpdatedAt: time.Unix(post.Post.UpdatedAt, 0),
		Title:     post.Post.Title,
		Description: sql.NullString{
			Valid:  true,
			String: post.Post.Description,
		},
		Pre:  post.Post.Pre,
		Next: post.Post.Next,
		Url:  post.Post.Url,
		Path: post.Post.Path,
		Tags: sql.NullString{
			Valid:  true,
			String: post.Post.Tags,
		},
	}

	return &types.Reply{
		Code: 666,
		Data: p,
	}, nil
}
