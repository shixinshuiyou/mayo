package main

import (
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/shixinshuiyou/mayo-dev/app/user/router"
	"github.com/shixinshuiyou/mayo-dev/config"
)

func main() {
	srvName := config.SrvActionName
	reg := etcdv3.NewRegistry(func(op *registry.Options) {

	})
	service := web.NewService(
		web.Name(srvName),
		web.Address(srvName),
		web.Handler(router.Register()),
		web.Registry(reg),
	)

	service.Init()
	service.Run()
	// if err := service.Run(); err != nil {

	// }
}
