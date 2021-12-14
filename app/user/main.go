package main

import (
	"fmt"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/shixinshuiyou/mayo/app/user/router"
	"github.com/shixinshuiyou/mayo/config"
)

func main() {
	srvName := config.SrvActionName
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"127.0.0.1:2380"}
	})
	service := web.NewService(
		web.Name(srvName),
		web.Handler(router.Register()),
		web.Registry(reg),
	)

	service.Init()
	r := router.Register()
	r.HandleMethodNotAllowed = true
	service.Handle("/", r)
	// Run server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
