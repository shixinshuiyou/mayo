package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/shixinshuiyou/mayo/config"
	proto "github.com/shixinshuiyou/mayo/proto/id"
	"github.com/shixinshuiyou/mayo/srv/id/snowflake"
	"github.com/shixinshuiyou/mayo/tool/log"
	"github.com/shixinshuiyou/mayo/tool/tracer"
)

func main() {
	srvName := config.SrvSnowflakeID
	log.InitLoggerJson(srvName)

	ot, closer, _ := tracer.SetJaegerGlobalTracer(srvName, config.JaegerAddress)
	defer closer.Close()

	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = config.EtcdAddress
	})

	service := micro.NewService(
		micro.Name(srvName),
		micro.Registry(reg),
		micro.WrapHandler(opentracing.NewHandlerWrapper(ot)),
	)

	service.Init()
	proto.RegisterIDHandler(service.Server(), new(snowflake.SnowID))
	// Register Struct as Subscriber
	// micro.RegisterSubscriber("go.micro.srv.hello", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	// micro.RegisterSubscriber("/snowflake", service.Server(), nil)

	// Run server
	if err := service.Run(); err != nil {
		log.Logger.Fatal(err)
	}
}
