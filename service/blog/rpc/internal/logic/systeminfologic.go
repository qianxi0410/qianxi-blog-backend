package logic

import (
	"context"

	"qianxi-blog/service/blog/rpc/blog"
	"qianxi-blog/service/blog/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSystemInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SystemInfoLogic {
	return &SystemInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SystemInfoLogic) SystemInfo(in *blog.SystemInfoReq) (*blog.SystemInfoReply, error) {
	// todo: add your logic here and delete this line

	return &blog.SystemInfoReply{}, nil
}
