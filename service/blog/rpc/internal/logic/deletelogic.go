package logic

import (
	"context"
	"qianxi-blog/common/key"

	"qianxi-blog/service/blog/rpc/blog"
	"qianxi-blog/service/blog/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLogic) Delete(in *blog.DeleteReq) (*blog.DeleteReply, error) {
	id := in.Id

	one, err := l.svcCtx.PostModel.FindOne(id)
	if err != nil {
		return nil, err
	}

	if one.Next != -1 {
		next, err := l.svcCtx.PostModel.FindOne(one.Next)
		if err != nil {
			return nil, err
		}
		next.Pre = one.Pre
		err = l.svcCtx.PostModel.Update(*next)
		if err != nil {
			return nil, err
		}
	}
	if one.Pre != -1 {
		pre, err := l.svcCtx.PostModel.FindOne(one.Pre)
		if err != nil {
			return nil, err
		}
		pre.Next = one.Next
		err = l.svcCtx.PostModel.Update(*pre)
		if err != nil {
			return nil, err
		}
	}

	err = l.svcCtx.PostModel.Delete(id)
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.CommentModel.DeleteByPostId(id)
	if err != nil {
		return nil, err
	}

	l.svcCtx.Redis.Expire(l.ctx, key.Post(id), 0)
	//l.svcCtx.Redis.Expire(l.ctx, key.Posts(page, size), 0)
	result, err := l.svcCtx.Redis.Keys(l.ctx, "*_*_page_size").Result()
	for i := 0; i < len(result); i++ {
		l.svcCtx.Redis.Expire(l.ctx, result[i], 0)
	}

	l.svcCtx.Redis.Expire(l.ctx, key.PostsCount(), 0)
	l.svcCtx.Redis.Expire(l.ctx, key.CountInfo(), 0)
	l.svcCtx.Redis.Expire(l.ctx, key.CommentCount(), 0)
	// TODO: 删除文件

	return &blog.DeleteReply{}, nil
}
