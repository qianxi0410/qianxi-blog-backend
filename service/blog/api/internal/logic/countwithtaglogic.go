package logic

import (
	"context"

	"qianxi-blog/service/blog/api/internal/svc"
	"qianxi-blog/service/blog/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CountWithTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCountWithTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) CountWithTagLogic {
	return CountWithTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CountWithTagLogic) CountWithTag(req types.CountWithTagReq) (*types.Reply, error) {
	// todo: add your logic here and delete this line

	return &types.Reply{}, nil
}
