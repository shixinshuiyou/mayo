package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/shixinshuiyou/framework/log"
	"github.com/shixinshuiyou/framework/tracer"
	"github.com/shixinshuiyou/mayo/app/user/router"
	"github.com/shixinshuiyou/mayo/config"
)

func main() {
	srvName := config.SrvActionName
	log.InitLoggerJson(srvName)

	jaegerTracer, closer, _ := tracer.InitJaegerTracer(srvName, "127.0.0.1:6831")
	// TODO 错误处理
	defer closer.Close()

	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"127.0.0.1:2380"}
	})

	service := web.NewService(
		web.Name(srvName),
		web.Handler(router.Register()),
		web.Registry(reg),
		web.MicroService(micro.NewService(micro.WrapHandler(opentracing.NewHandlerWrapper(jaegerTracer)))),
	)

	log.Logger.Debugf("")

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
