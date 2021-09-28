package main

import (
	"flag"
	"fmt"

	"github.com/tal-tech/go-zero/rest/httpx"

	"qianxi-blog/service/blog/api/internal/config"
	"qianxi-blog/service/blog/api/internal/handler"
	"qianxi-blog/service/blog/api/internal/svc"

	errorHandler "qianxi-blog/common/handler"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/post-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	httpx.SetErrorHandler(errorHandler.ReturnErrorHandler)
	handler.RegisterHandlers(server, ctx)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
