package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/qianxi/blog-backend/controller"
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
	r.Use(cors.Default())
	addr, port := viper.GetString("server.address"), viper.GetInt("server.port")

	post := r.Group("/post")
	{
		var pc controller.PostController
		post.GET("/:id", pc.GetPostById)
		post.GET("/page/:page/size/:size", pc.GetPostByPageQuery)
		post.GET("/page/:page/size/:size/tag/:tag", pc.GetPostByPageAndTagQuery)
		post.GET("/count", pc.GetCount)
		post.GET("/count/tag/:tag", pc.GetCountWithTag)
	}

	oauth2 := r.Group("/oauth2")
	{
		var oc controller.OAuth2Controller
		oauth2.GET("/code/:code", oc.GetUserInfo)
	}

	if err := r.Run(fmt.Sprintf("%s:%d", addr, port)); err != nil {
		log.Fatalf("start server failed : %v", err)
	}
}
