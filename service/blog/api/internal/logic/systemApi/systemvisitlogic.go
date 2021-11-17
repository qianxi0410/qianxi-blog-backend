package systemApi

import (
	"context"
	"net/http"
	"qianxi-blog/service/blog/model"
	"time"

	"qianxi-blog/service/blog/api/internal/svc"
	"qianxi-blog/service/blog/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemVisitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemVisitLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemVisitLogic {
	return SystemVisitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemVisitLogic) SystemVisit(req types.VisitReq, r *http.Request) (*types.Reply, error) {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = r.RemoteAddr
	}

	_, err := l.svcCtx.VisitModel.Insert(model.Visit{
		VisitTime: time.Now(),
		Ip:        ip,
		Uri:       req.Uri,
	})
	if err != nil {
		return nil, err
	}

	return &types.Reply{
		Code: 666,
	}, nil
}
