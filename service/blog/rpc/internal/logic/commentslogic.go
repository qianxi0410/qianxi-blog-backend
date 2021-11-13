package logic

import (
	"context"
	"encoding/json"
	"qianxi-blog/common/key"
	"qianxi-blog/service/blog/model"
	"qianxi-blog/service/blog/rpc/blog"
	"qianxi-blog/service/blog/rpc/blogclient"
	"qianxi-blog/service/blog/rpc/internal/svc"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/tal-tech/go-zero/core/logx"
)

type CommentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentsLogic {
	return &CommentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentsLogic) Comments(in *blog.PageReq) (*blog.CommentPageReply, error) {
	page, size := in.Page, in.Size

	if size == -1 {
		all, err := l.svcCtx.CommentModel.CommentsAll()
		if err != nil {
			return nil, err
		}

		var comments []*blogclient.Comment
		for i := 0; i < len(all); i++ {
			comments = append(comments, &blogclient.Comment{
				Id:         all[i].Id,
				CreatedAt:  all[i].CreatedAt.Unix(),
				UpdateedAt: all[i].UpdateedAt.Unix(),
				Content:    all[i].Content,
				Login:      all[i].Login,
				Name:       all[i].Name,
				Avatar:     all[i].Avatar,
				PostId:     all[i].PostId,
			})
		}

		return &blogclient.CommentPageReply{Comments: comments}, nil
	}

	var comments []model.Comments

	bytes, err := l.svcCtx.Redis.Get(l.ctx, key.Comments(page, size)).Bytes()
	if err != redis.Nil {
		err := json.Unmarshal(bytes, &comments)
		if err != nil {
			return nil, err
		}
	} else {
		offset := (page - 1) * size

		comments, err = l.svcCtx.CommentModel.Comments(size, offset)
		if err != nil {
			return nil, err
		}

		marshal, err := json.Marshal(comments)
		if err != nil {
			return nil, err
		}

		l.svcCtx.Redis.Set(l.ctx, key.Comments(page, size), marshal, 10*time.Minute)
	}

	var reply []*blogclient.Comment
	for i := 0; i < len(comments); i++ {
		reply = append(reply, &blogclient.Comment{
			Id:         comments[i].Id,
			CreatedAt:  comments[i].CreatedAt.Unix(),
			UpdateedAt: comments[i].UpdateedAt.Unix(),
			Content:    comments[i].Content,
			Login:      comments[i].Login,
			Name:       comments[i].Name,
			Avatar:     comments[i].Avatar,
			PostId:     comments[i].PostId,
		})
	}

	return &blog.CommentPageReply{
		Comments: reply,
	}, nil
}
