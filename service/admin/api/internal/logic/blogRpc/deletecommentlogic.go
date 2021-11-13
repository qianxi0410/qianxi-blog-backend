package blogRpc

import (
	"context"
	"qianxi-blog/service/blog/rpc/blogclient"

	"qianxi-blog/service/admin/api/internal/svc"
	"qianxi-blog/service/admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteCommentLogic {
	return DeleteCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCommentLogic) DeleteComment(req types.DeleteCommentReq) (*types.Reply, error) {

	_, err := l.svcCtx.BlogRpc.DeleteComment(l.ctx, &blogclient.CommentDeleteReq{
		Id:     req.Id,
		PostId: req.PostId,
	})

	if err != nil {
		return nil, err
	}

	return &types.Reply{
		Code: 666,
	}, nil
}
