package main

import (
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/micro/micro/v2/cmd"
	"github.com/micro/micro/v2/plugin"
	"github.com/shixinshuiyou/mayo/config"
	"github.com/shixinshuiyou/mayo/tool/auth"
	"github.com/shixinshuiyou/mayo/tool/hystrix"
	"github.com/shixinshuiyou/mayo/tool/tracer"
)

func init() {
	// 用户鉴权
	plugin.Register(plugin.NewPlugin(
		plugin.WithName("auth"),
		plugin.WithHandler(auth.APiGatewayAuth),
	))

	// 配置断路器
	plugin.Register(plugin.NewPlugin(
		plugin.WithName("hystrix"),
		plugin.WithHandler(hystrix.HystrixHandler),
	))

	plugin.Register(plugin.NewPlugin(
		plugin.WithName("tracer"),
		plugin.WithHandler(tracer.TracerWrapper),
	))
}

func main() {
	srvName := config.SrvApiGateAway
	_, closer, _ := tracer.SetJaegerGlobalTracer(srvName, config.JaegerAddress)
	defer closer.Close()

	cmd.Init(
		micro.Name(srvName),
		micro.Registry(etcdv3.NewRegistry(func(op *registry.Options) {
			op.Addrs = config.EtcdAddress
		})),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
}
