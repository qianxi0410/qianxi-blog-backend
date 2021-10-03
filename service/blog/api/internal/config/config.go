package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	// rest url
	rest.RestConf
	// $user:$password@tcp($ip:$port)/$db?$queries
	Mysql struct {
		DataSource string
	}
	// cache
	CacheRedis cache.CacheConf

	// redis
	Redis redis.RedisConf

	// github
	Github struct {
		ApiUrl       string
		TokenUrl     string
		ClientId     string
		ClientSecret string
	}
}
