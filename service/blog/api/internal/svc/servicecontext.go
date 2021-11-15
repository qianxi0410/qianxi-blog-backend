package svc

import (
	"qianxi-blog/service/blog/api/internal/config"
	"qianxi-blog/service/blog/model"

	"github.com/go-redis/redis/v8"

	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	PostModel    model.PostsModel
	CommentModel model.CommentsModel
	SystemModel  model.SystemModel
	Redis        *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:       c,
		PostModel:    model.NewPostsModel(conn),
		CommentModel: model.NewCommentsModel(conn),
		SystemModel:  model.NewSystemModel(conn),
		Redis: redis.NewClient(&redis.Options{
			Addr:     c.Redis.Host,
			Password: c.Redis.Pass,
		}),
	}
}
