package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/qianxi/blog_backend/controller"
	"github.com/spf13/viper"
)

const configPath = "./config/config.json"

func init() {
	// 读取配置文件
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
}

func main() {
	r := gin.Default()

	addr, port := viper.GetString("server.address"), viper.GetInt("server.port")

	post := r.Group("/post")
	{
		var pc controller.PostController
		post.GET("/:id", pc.GetPostById)
	}

	if err := r.Run(fmt.Sprintf("%s:%d", addr, port)); err != nil {
		panic("start server failed !")
	}
}
