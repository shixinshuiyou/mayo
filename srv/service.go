package srv

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/shixinshuiyou/mayo/config"
	proto "github.com/shixinshuiyou/mayo/proto/id"
)

var idService proto.IDService

// 统一管理gRPC服务调用
func init() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = config.EtcdAddress
	})
	srv := micro.NewService(micro.Registry(reg))
	idService = proto.NewIDService(config.SrvSnowflakeID, srv.Client())
}
