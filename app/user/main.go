package main

import (
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/shixinshuiyou/mayo/app/user/router"
	"github.com/shixinshuiyou/mayo/config"
	"github.com/shixinshuiyou/mayo/tool/log"
	"github.com/shixinshuiyou/mayo/tool/tracer"
	_ "github.com/shixinshuiyou/mayo/tool/validator"
)

func main() {
	srvName := config.SrvActionName
	log.InitLoggerJson(srvName)

	_, closer, _ := tracer.SetJaegerGlobalTracer(srvName, config.JaegerAddress)
	defer closer.Close()

	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = config.EtcdAddress
	})

	service := web.NewService(
		web.Name(srvName),
		web.Handler(router.Register()),
		web.Registry(reg),
	)

	if err := service.Init(); err != nil {
		log.Logger.Errorf("servce(%s) init error:%s", srvName, err)
	}

	r := router.Register()
	r.HandleMethodNotAllowed = true
	service.Handle("/", r)
	// Run server
	if err := service.Run(); err != nil {
		log.Logger.Errorf("servce(%s) run error:%s", srvName, err)
	}
}
