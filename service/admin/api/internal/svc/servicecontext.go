package svc

import (
	"qianxi-blog/service/admin/api/internal/config"
	"qianxi-blog/service/admin/model"
	"qianxi-blog/service/blog/rpc/blogclient"

	"github.com/tal-tech/go-zero/zrpc"

	"github.com/tal-tech/go-zero/core/stores/sqlx"

	"github.com/go-redis/redis/v8"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UsersModel
	Redis     *redis.Client
	BlogRpc   blogclient.Blog
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
		BlogRpc: blogclient.NewBlog(zrpc.MustNewClient(c.BlogRpc)),
	}
}
