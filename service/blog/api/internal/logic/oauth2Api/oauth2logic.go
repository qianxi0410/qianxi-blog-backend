package oauth2Api

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"qianxi-blog/service/blog/api/internal/svc"
	"qianxi-blog/service/blog/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type Oauth2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOauth2Logic(ctx context.Context, svcCtx *svc.ServiceContext) Oauth2Logic {
	return Oauth2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Oauth2Logic) Oauth2(req types.Oauth2Req) (*types.Reply, error) {
	resp, err := http.Post(fmt.Sprintf("%s?client_id=%s&client_secret=%s&code=%s",
		l.svcCtx.Config.Github.TokenUrl, l.svcCtx.Config.Github.ClientId,
		l.svcCtx.Config.Github.ClientSecret, req.Code), "application/json;charset=utf-8", nil)

	if err != nil {
		return nil, errors.New("认证失败: " + err.Error())
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, errors.New("验证失败: " + err.Error())
	}

	accessToken := strings.Split(strings.Split(string(body), "&")[0], "=")[1]

	httpReq, err := http.NewRequest(http.MethodGet, l.svcCtx.Config.Github.ApiUrl, nil)
	if err != nil {
		return nil, errors.New("验证失败: " + err.Error())
	}

	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Authorization", fmt.Sprintf("token %s", accessToken))

	githubResp, err := (&http.Client{}).Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer githubResp.Body.Close()

	info, err := io.ReadAll(githubResp.Body)

	if err != nil {
		return nil, errors.New("验证失败: " + err.Error())
	}

	s := string(info)

	return &types.Reply{
		Code: 666,
		Data: s,
	}, nil
}
