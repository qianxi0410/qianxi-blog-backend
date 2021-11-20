package main

import (
	"flag"
	"fmt"
	errorHandler "qianxi-blog/common/handler"

	"github.com/tal-tech/go-zero/rest/httpx"

	"qianxi-blog/service/admin/api/internal/config"
	"qianxi-blog/service/admin/api/internal/handler"
	"qianxi-blog/service/admin/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/admin-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	//server := rest.MustNewServer(c.RestConf, rest.WithNotAllowedHandler(middlewares.NewCorsMiddleware().Handler()))
	defer server.Stop()
	//server.Use(middlewares.NewCorsMiddleware().Handle)

	httpx.SetErrorHandler(errorHandler.ReturnErrorHandler)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
