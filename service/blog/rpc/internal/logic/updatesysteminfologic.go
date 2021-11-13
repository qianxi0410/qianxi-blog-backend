package logic

import (
	"context"
	"qianxi-blog/common/key"

	"qianxi-blog/service/blog/rpc/blog"
	"qianxi-blog/service/blog/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateSystemInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSystemInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSystemInfoLogic {
	return &UpdateSystemInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSystemInfoLogic) UpdateSystemInfo(in *blog.UpdateSystemInfoReq) (*blog.UpdateSystemInfoReply, error) {
	k, value := in.Key, in.Value

	one, err := l.svcCtx.SystemModel.FindOne(k)
	if err != nil {
		return nil, err
	}

	one.Value = value
	err = l.svcCtx.SystemModel.Update(*one)
	if err != nil {
		return nil, err
	}

	l.svcCtx.Redis.Expire(l.ctx, key.AllInfo(), 0)

	return &blog.UpdateSystemInfoReply{}, nil
}
