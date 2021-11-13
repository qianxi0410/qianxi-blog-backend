package blogRpc

import (
	"context"
	"errors"
	"qianxi-blog/service/blog/model"
	"qianxi-blog/service/blog/rpc/blogclient"
	"time"

	"qianxi-blog/service/admin/api/internal/svc"
	"qianxi-blog/service/admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CommentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) CommentsLogic {
	return CommentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentsLogic) Comments(req types.PageReq) (*types.Reply, error) {
	page, size := req.Page, req.Size

	if (page <= 0 || size < 0) && size != -1 {
		return nil, errors.New("分页参数不合法")
	}

	comments, err := l.svcCtx.BlogRpc.Comments(l.ctx, &blogclient.PageReq{
		Page: page,
		Size: size,
	})

	if err != nil {
		return nil, err
	}

	var reply []model.Comments

	for i := 0; i < len(comments.Comments); i++ {
		reply = append(reply, model.Comments{
			Id:         comments.Comments[i].Id,
			CreatedAt:  time.Unix(comments.Comments[i].CreatedAt, 0),
			UpdateedAt: time.Unix(comments.Comments[i].UpdateedAt, 0),
			Content:    comments.Comments[i].Content,
			Login:      comments.Comments[i].Login,
			Name:       comments.Comments[i].Name,
			Avatar:     comments.Comments[i].Avatar,
			PostId:     comments.Comments[i].PostId,
		})
	}

	return &types.Reply{
		Code: 666,
		Data: reply,
	}, nil
}
