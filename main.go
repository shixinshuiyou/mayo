package main

import (
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/micro/micro/v2/cmd"
	"github.com/shixinshuiyou/framework/tracer"
	"github.com/shixinshuiyou/mayo/config"
)

func main() {
	srvName := config.SrvApiGateAway
	jaegerTracer, closer, _ := tracer.InitJaegerTracer(srvName, "127.0.0.1:6381")
	// TODO 错误处理
	defer closer.Close()

	cmd.Init(
		micro.Name(srvName),
		micro.Registry(etcdv3.NewRegistry(func(op *registry.Options) {

		})),
		micro.WrapHandler(opentracing.NewHandlerWrapper(jaegerTracer)),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Address("127.0.0.1:8086"),
	)
}
