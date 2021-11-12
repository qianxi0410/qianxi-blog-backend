package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	// $user:$password@tcp($ip:$port)/$db?$queries
	Mysql struct {
		DataSource string
	}
	// redis
	Redis redis.RedisConf

	// cache
	CacheRedis cache.CacheConf
}
