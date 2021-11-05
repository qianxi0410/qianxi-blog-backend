package userApi

import (
	"context"
	"errors"

	"qianxi-blog/service/admin/api/internal/svc"
	"qianxi-blog/service/admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (*types.Reply, error) {
	username, password := req.Name, req.Password

	user, err := l.svcCtx.UserModel.FindWithNameAndPass(username, password)

	if err != nil {
		return nil, errors.New("用户名或密码错误: " + err.Error())
	}

	if user == nil {
		return nil, errors.New("用户名或密码错误")
	}

	return &types.Reply{
		Code: 666,
		Data: true,
	}, nil
}
