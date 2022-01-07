package main

import (
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/shixinshuiyou/mayo/config"
	"github.com/shixinshuiyou/mayo/proto"
	"github.com/shixinshuiyou/mayo/tool/log"
)

func main() {
	srvName := config.SrvSnowflakeID
	log.InitLoggerJson(srvName)

	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = config.EtcdAddress
	})

	service := web.NewService(
		web.Name(srvName),
		web.Registry(reg),
	)

	service.Init()
	proto.RegisterIDHandler(server.NewServer(), nil)

	// Run server
	if err := service.Run(); err != nil {
		log.Logger.Errorf("servce(%s) run error:%s", srvName, err)
	}
}
