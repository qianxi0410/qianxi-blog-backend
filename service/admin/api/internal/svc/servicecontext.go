package svc

import (
	"qianxi-blog/service/admin/api/internal/config"
	"qianxi-blog/service/admin/model"

	"github.com/tal-tech/go-zero/core/stores/sqlx"

	"github.com/go-redis/redis/v8"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UsersModel
	Redis     *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUsersModel(conn),
		Redis: redis.NewClient(&redis.Options{
			Addr:     c.Redis.Host,
			Password: c.Redis.Pass,
		}),
	}
}
