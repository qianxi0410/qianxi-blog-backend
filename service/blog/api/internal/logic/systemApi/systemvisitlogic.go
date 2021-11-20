package systemApi

import (
	"context"
	"net/http"
	"qianxi-blog/service/blog/model"
	"strings"
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

func getClientIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if strings.Contains(ip, "127.0.0.1") || ip == "" {
		ip = r.Header.Get("X-Real-IP")
	}

	if ip == "" {
		return "127.0.0.1"
	}

	return ip
}

func (l *SystemVisitLogic) SystemVisit(req types.VisitReq, r *http.Request) (*types.Reply, error) {
	ip := getClientIP(r)

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
