package blogRpc

import (
	"context"
	"qianxi-blog/service/blog/rpc/blog"

	"qianxi-blog/service/admin/api/internal/svc"
	"qianxi-blog/service/admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type InsertPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) InsertPostLogic {
	return InsertPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertPostLogic) InsertPost(req types.InsertReq) (*types.Reply, error) {
	_, err := l.svcCtx.BlogRpc.Insert(l.ctx, &blog.InsertReq{
		Title:       req.Title,
		Description: req.Description,
		Url:         req.Url,
		Path:        req.Path,
		Tags:        req.Tags,
		Blur:        req.Blur,
	})

	if err != nil {
		return nil, err
	}

	return &types.Reply{
		Code: 666,
	}, nil
}
