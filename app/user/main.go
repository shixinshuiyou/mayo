package main

import (
	"fmt"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/shixinshuiyou/mayo-dev/app/user/router"
)

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {

	})
	service := web.NewService(
		web.Name("czh.micro.api.action"),
		web.Address("127.0.0.1:9000"),
		web.Handler(router.Register()),
		web.Registry(reg),
	)

	service.Init()

	if err := service.Run(); err != nil {
		fmt.Printf("error:%s", err)
	}
}
