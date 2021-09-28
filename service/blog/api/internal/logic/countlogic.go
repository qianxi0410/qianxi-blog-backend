package logic

import (
	"context"
	"errors"

	"qianxi-blog/service/blog/api/internal/svc"
	"qianxi-blog/service/blog/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) CountLogic {
	return CountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CountLogic) Count() (*types.Reply, error) {
	res, err := l.svcCtx.PostModel.Count()

	if err != nil {
		return nil, errors.New("查询博客总数出错")
	}

	return &types.Reply{
		Code: 777,
		Data: res,
	}, nil
}
