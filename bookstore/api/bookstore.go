package main

import (
	"flag"
	"fmt"

	"bookstore/api/internal/config"
	"bookstore/api/internal/handler"
	"bookstore/api/internal/svc"

	"github.com/sllt/tao/core/conf"
	"github.com/sllt/tao/rest"
)

var configFile = flag.String("f", "etc/bookstore-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
