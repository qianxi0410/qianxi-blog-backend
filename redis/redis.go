package redis

import (
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var options redis.Options

func init() {
	viper.SetConfigFile("./config/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}

	options.Addr = viper.GetString("redis.addr")
	options.Password = viper.GetString("redis.password")
	options.DB = 0
}

func New() *redis.Client {
	return redis.NewClient(&options)
}
