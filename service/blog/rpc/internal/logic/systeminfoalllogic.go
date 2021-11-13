package logic

import (
	"context"
	"qianxi-blog/common/key"

	"qianxi-blog/service/blog/rpc/blog"
	"qianxi-blog/service/blog/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemInfoAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSystemInfoAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SystemInfoAllLogic {
	return &SystemInfoAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SystemInfoAllLogic) SystemInfoAll(in *blog.SystemInfoAllReq) (*blog.SystemInfoAllReply, error) {
	m, err := l.svcCtx.Redis.HGetAll(l.ctx, key.AllInfo()).Result()
	if err != nil {
		return nil, err
	}

	if len(m) != 0 {
		return &blog.SystemInfoAllReply{Kv: m}, nil
	}

	all, err := l.svcCtx.SystemModel.All()
	if err != nil {
		return nil, err
	}

	m = make(map[string]string, len(all))

	for i := 0; i < len(all); i++ {
		m[all[i].Key] = all[i].Value
	}

	l.svcCtx.Redis.HSet(l.ctx, key.AllInfo(), m)

	return &blog.SystemInfoAllReply{
		Kv: m,
	}, nil
}
