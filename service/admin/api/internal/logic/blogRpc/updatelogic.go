package blogRpc

import (
	"context"
	"qianxi-blog/service/admin/api/internal/svc"
	"qianxi-blog/service/admin/api/internal/types"
	"qianxi-blog/service/blog/rpc/blog"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateLogic {
	return UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req types.UpdateReq) (*types.Reply, error) {
	_, err := l.svcCtx.BlogRpc.Update(l.ctx, &blog.UpdateReq{Post: &blog.Post{
		Id:          req.Id,
		CreatedAt:   req.CreatedAt,
		Title:       req.Title,
		Description: req.Description,
		Pre:         req.Pre,
		Next:        req.Next,
		Url:         req.Url,
		Path:        req.Path,
		Tags:        req.Tags,
		Blur:        req.Blur,
	}, Page: req.Page, Size: req.Size})

	if err != nil {
		return nil, err
	}

	return &types.Reply{
		Code: 666,
		Data: nil,
	}, nil
}
