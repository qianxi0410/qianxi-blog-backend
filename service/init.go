package service

import (
	"context"

	"github.com/go-redis/redis/v8"
	rdx "github.com/qianxi/blog-backend/redis"
)

var rdb *redis.Client
var ctx context.Context

func init() {
	rdb = rdx.New()
	ctx = context.Background()
}
