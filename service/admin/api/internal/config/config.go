package config

import (
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
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
		AccessSecret string
		AccessExpire int64
		Issuer       string
	}

	BlogRpc zrpc.RpcClientConf
}
