package userApi

import (
	"context"

	"qianxi-blog/service/admin/api/internal/svc"
	"qianxi-blog/service/admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginValidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginValidLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginValidLogic {
	return LoginValidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginValidLogic) LoginValid() (*types.Reply, error) {
	return &types.Reply{
		Code: 666,
		Data: true,
	}, nil
}
