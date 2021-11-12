package blogRpc

import (
	"context"
	"errors"
	"qianxi-blog/service/blog/rpc/blogclient"

	"qianxi-blog/service/admin/api/internal/svc"
	"qianxi-blog/service/admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CommentCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) CommentCountLogic {
	return CommentCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentCountLogic) CommentCount() (*types.Reply, error) {
	count, err := l.svcCtx.BlogRpc.CommentCount(l.ctx, &blogclient.CountReq{})
	if err != nil {
		return nil, errors.New("查询文章总数时出错:" + err.Error())
	}

	return &types.Reply{
		Code: 666,
		Data: count.Count,
	}, nil

}
