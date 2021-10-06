package commentApi

import (
	"context"
	"errors"
	"fmt"
	"qianxi-blog/common/key"
	"qianxi-blog/service/blog/model"
	"time"

	"qianxi-blog/service/blog/api/internal/svc"
	"qianxi-blog/service/blog/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) SaveLogic {
	return SaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveLogic) Save(req types.SaveReq) (*types.Reply, error) {
	if len(req.Content) == 0 || len(req.Login) == 0 {
		return nil, errors.New("评论失败: 内容或登陆名字段为空")
	}

	if req.PostId <= 0 {
		return nil, errors.New("评论失败: 评论文章不存在")
	}

	var commnet = model.Comments{
		CreatedAt:  time.Now(),
		UpdateedAt: time.Now(),
		Content:    req.Content,
		Login:      req.Login,
		Name:       req.Name,
		Avatar:     req.Avatar,
		PostId:     req.PostId,
	}
	ret, err := l.svcCtx.CommentModel.Insert(commnet)

	if err != nil {
		return nil, errors.New("评论失败: " + err.Error())
	}

	l.svcCtx.Redis.ExpireAt(context.Background(), key.Post(req.PostId), time.Now())

	id, err := ret.LastInsertId()
	fmt.Println(id)
	if err != nil {
		return nil, errors.New("评论失败: " + err.Error())
	}

	return &types.Reply{
		Code: 666,
		Data: id,
	}, nil
}
