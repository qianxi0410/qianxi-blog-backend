package blogRpc

import (
	"context"
	"database/sql"
	"errors"
	"qianxi-blog/service/blog/model"
	"qianxi-blog/service/blog/rpc/blogclient"
	"time"

	"qianxi-blog/service/admin/api/internal/svc"
	"qianxi-blog/service/admin/api/internal/types"

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
	page, size := req.Page, req.Size

	if (page <= 0 || size <= 0) && size != -1 {
		return nil, errors.New("分页查询参数错误")
	}

	posts, err := l.svcCtx.BlogRpc.Posts(l.ctx, &blogclient.PageReq{
		Page: page,
		Size: size,
	})

	if err != nil {
		return nil, errors.New("远程调用错误: " + err.Error())
	}

	var reply []*model.Posts

	for i := 0; i < len(posts.Posts); i++ {
		reply = append(reply, &model.Posts{
			Id:        posts.Posts[i].Id,
			CreatedAt: time.Unix(posts.Posts[i].CreatedAt, 0),
			UpdatedAt: time.Unix(posts.Posts[i].UpdatedAt, 0),
			Title:     posts.Posts[i].Title,
			Description: sql.NullString{
				String: posts.Posts[i].Description,
				Valid:  true,
			},
			Pre:  posts.Posts[i].Pre,
			Next: posts.Posts[i].Next,
			Url:  posts.Posts[i].Url,
			Path: posts.Posts[i].Path,
			Tags: sql.NullString{
				String: posts.Posts[i].Tags,
				Valid:  true,
			},
			Blur: posts.Posts[i].Blur,
		})
	}

	return &types.Reply{
		Code: 666,
		Data: reply,
	}, nil
}
