package commentApi

import (
	"context"
	"errors"
	"qianxi-blog/common/key"
	"time"

	"qianxi-blog/service/blog/api/internal/svc"
	"qianxi-blog/service/blog/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteLogic {
	return DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req types.DeleteReq) (*types.Reply, error) {
	if req.Id <= 0 {
		return nil, errors.New("删除评论失败: 输入的评论id有误")
	}

	comment, err := l.svcCtx.CommentModel.FindOne(req.Id)
	if err != nil {
		return nil, errors.New("删除评论失败: " + err.Error())
	}

	if len(req.Login) == 0 || req.Login != comment.Login {
		return nil, errors.New("你没有权限删除本条评论")
	}

	l.svcCtx.Redis.ExpireAt(context.Background(), key.Post(comment.PostId), time.Now())

	err = l.svcCtx.CommentModel.Delete(req.Id)
	if err != nil {
		return nil, errors.New("删除评论失败: " + err.Error())
	}

	return &types.Reply{
		Code: 666,
	}, nil
}
