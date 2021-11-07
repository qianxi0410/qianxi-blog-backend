package userApi

import (
	"context"
	"errors"
	"qianxi-blog/service/admin/api/internal/logic/utils"
	"time"

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

	tokenStr, err := utils.GenerateToken(utils.JwtInfo{
		Issuer: l.svcCtx.Config.Jwt.Issuer,
		Secret: l.svcCtx.Config.Jwt.AccessSecret,
	}, user.Name, time.Minute*time.Duration(l.svcCtx.Config.Jwt.AccessExpire))

	if err != nil {
		return nil, errors.New("签发token出错: " + err.Error())
	}

	return &types.Reply{
		Code: 666,
		Data: tokenStr,
	}, nil
}
