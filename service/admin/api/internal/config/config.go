package config

import (
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	rest.RestConf

	// $user:$password@tcp($ip:$port)/$db?$queries
	Mysql struct {
		DataSource string
	}

	// redis
	Redis redis.RedisConf

	Jwt struct {
		Secret string
		Issuer string
	}
}
