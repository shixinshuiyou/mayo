package config

import (
	"strings"

	"github.com/shixinshuiyou/mayo/tool/config"
)

const (
	SrvApiGateAway = "czh.micro.api"
	SrvActionName  = "czh.micro.api.user"
	SrvSnowflakeID = "czh.micro.srv.id"
)

var (
	JaegerAddress string
	EtcdAddress   []string
	// Once          sync.Once
)

// 读取配置文件
func init() {
	JaegerAddress = config.Conf.Get("jaeger", "address").String("127.0.0.1:6831")
	EtcdAddr := config.Conf.Get("etcd", "address").String("127.0.0.1:2379")
	EtcdAddress = strings.Split(EtcdAddr, ",")
}
