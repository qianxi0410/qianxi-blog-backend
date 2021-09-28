package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"qianxi-blog/service/blog/api/internal/config"
	"qianxi-blog/service/blog/model"
)

type ServiceContext struct {
	Config config.Config
	PostModel 	model.PostsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config: c,
		PostModel: model.NewPostsModel(conn),
	}
}
