package systemApi

import (
	"context"
	"qianxi-blog/common/key"
	"qianxi-blog/service/blog/api/internal/svc"
	"qianxi-blog/service/blog/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemInfoLogic {
	return SystemInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemInfoLogic) SystemInfo() (*types.Reply, error) {
	m, err := l.svcCtx.Redis.HGetAll(l.ctx, key.AllInfo()).Result()
	if err != nil {
		return nil, err
	}

	if len(m) != 0 {
		return &types.Reply{
			Code: 666,
			Data: m,
		}, nil
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

	return &types.Reply{
		Code: 666,
		Data: m,
	}, nil
}
