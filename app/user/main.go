package main

import (
	"fmt"

	"github.com/micro/go-micro/v2/web"
	"github.com/shixinshuiyou/mayo-dev/app/user/router"
)

func main() {
	service := web.NewService(
		web.Name("czh.micro.api.action"),
		web.Address("127.0.0.1:9000"),
		web.Handler(router.Register()),
	)

	service.Init()

	if err := service.Run(); err != nil {
		fmt.Printf("error:%s", err)
	}
}
