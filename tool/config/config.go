package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/etcd"
	"github.com/micro/go-micro/v2/config/source/file"
)

var (
	// Conf 采用micro-conf 配置中心
	Conf config.Config

	defaultPrefix = "/mayo/config"
)

func init() {
	Conf, _ = config.NewConfig()

	// 加载和合并多个源。合并优先级顺序相反. 此处  etcd > file
	// 官方支持的解析器还有 yaml、toml、xml、hcl
	Conf.Load(
		file.NewSource(file.WithPath(GetConfigFilePath())),
		etcd.NewSource(
			etcd.WithAddress(GetEtcdAddr()...),
			etcd.WithPrefix(defaultPrefix),
			etcd.StripPrefix(true),
		),
	)

	Conf.Sync()

}

func GetConfigFilePath() string {
	if GetMode() == "dev" {
		return "/Users/shixinshuiyou/go/mayo/docker/dev/conf.yaml"
	}
	return fmt.Sprintf("docker/%s/conf.yaml", GetMode())
}

func GetEtcdAddr() []string {
	addr := os.Getenv("ETCD_ADDR")
	return strings.Split(addr, ",")
}

func GetMode() string {
	env := os.Getenv("RUN_MODE")
	//dev,prod,test
	if env == "" {
		env = "dev"
	}
	return env
}
