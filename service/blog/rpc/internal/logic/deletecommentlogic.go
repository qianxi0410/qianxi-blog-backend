package logic

import (
	"context"
	"qianxi-blog/common/key"

	"qianxi-blog/service/blog/rpc/blog"
	"qianxi-blog/service/blog/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCommentLogic) DeleteComment(in *blog.CommentDeleteReq) (*blog.CommentDeleteReply, error) {
	commentId, postId := in.Id, in.PostId

	err := l.svcCtx.CommentModel.Delete(commentId)
	if err != nil {
		return nil, err
	}

	l.svcCtx.Redis.Expire(l.ctx, key.Post(postId), 0)
	l.svcCtx.Redis.Expire(l.ctx, key.CommentCount(), 0)
	l.svcCtx.Redis.Expire(l.ctx, key.CountInfo(), 0)
	result, err := l.svcCtx.Redis.Keys(l.ctx, "comment:*_*_page_size").Result()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(result); i++ {
		l.svcCtx.Redis.Expire(l.ctx, result[i], 0)
	}

	return &blog.CommentDeleteReply{}, nil
}
