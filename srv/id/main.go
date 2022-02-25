package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/shixinshuiyou/mayo/config"
	proto "github.com/shixinshuiyou/mayo/proto/id"
	"github.com/shixinshuiyou/mayo/srv/id/snowflake"
	"github.com/shixinshuiyou/mayo/tool/log"
	"github.com/shixinshuiyou/mayo/tool/tracer"
)

func main() {
	srvName := config.SrvSnowflakeID
	log.InitLoggerJson(srvName)

	_, closer, _ := tracer.SetJaegerGlobalTracer(srvName, config.JaegerAddress)
	defer closer.Close()

	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"127.0.0.1:2379"}
	})

	service := micro.NewService(
		micro.Name(srvName),
		micro.Registry(reg),
		micro.WrapHandler(tracer.NewHandlerWrapper()),
	)

	service.Init()
	proto.RegisterIDHandler(service.Server(), new(snowflake.SnowID))

	// Run server
	if err := service.Run(); err != nil {
		log.Logger.Errorf("servce(%s) run error:%s", srvName, err)
	}
}
