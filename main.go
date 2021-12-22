package main

import (
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/micro/micro/v2/cmd"
	"github.com/micro/micro/v2/plugin"
	"github.com/shixinshuiyou/framework/tracer"
	"github.com/shixinshuiyou/mayo/config"
	"github.com/shixinshuiyou/mayo/tool/auth"
)

func init() {
	plugin.Register(plugin.NewPlugin(
		plugin.WithName("auth"),
		plugin.WithHandler(auth.APiGatewayAuth),
	))
}

func main() {
	srvName := config.SrvApiGateAway
	jaegerTracer, closer, _ := tracer.InitJaegerTracer(srvName, config.JaegerAddress)
	// TODO 错误处理
	defer closer.Close()

	cmd.Init(
		micro.Name(srvName),
		micro.Registry(etcdv3.NewRegistry(func(op *registry.Options) {
			op.Addrs = []string{config.EtcdAddress}
		})),
		micro.WrapHandler(opentracing.NewHandlerWrapper(jaegerTracer)),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

}
