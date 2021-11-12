package blogRpc

import (
	"context"
	"qianxi-blog/service/blog/rpc/blog"

	"qianxi-blog/service/admin/api/internal/svc"
	"qianxi-blog/service/admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdatWithContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatWithContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdatWithContentLogic {
	return UpdatWithContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatWithContentLogic) UpdatWithContent(req types.UpdateReq) (*types.Reply, error) {
	_, err := l.svcCtx.BlogRpc.UpdateWithContent(l.ctx, &blog.UpdateReq{Post: &blog.Post{
		Id:          req.Id,
		CreatedAt:   req.CreatedAt,
		Title:       req.Title,
		Description: req.Description,
		Pre:         req.Pre,
		Next:        req.Next,
		Url:         req.Url,
		Path:        req.Path,
		Tags:        req.Tags,
	}, Page: req.Page, Size: req.Size})

	if err != nil {
		return nil, err
	}

	return &types.Reply{
		Code: 666,
		Data: nil,
	}, nil
}
