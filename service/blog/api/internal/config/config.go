package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	// rest url
	rest.RestConf
	// $user:$password@tcp($ip:$port)/$db?$queries
	Mysql struct {
		DataSource string
	}
	// redis
	CacheRedis cache.CacheConf
	// rpc client
	UserRpc    zrpc.RpcClientConf
}
