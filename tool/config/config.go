package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
	"github.com/shixinshuiyou/mayo/tool/config/etcd"
	"github.com/shixinshuiyou/mayo/tool/log"
)

var (
	// Conf 采用micro-conf 配置中心
	Conf config.Config
)

func init() {
	Conf, _ = config.NewConfig()

	// 加载和合并多个源。合并优先级顺序相反. 此处  etcd > file
	// 官方支持的解析器还有 yaml、toml、xml、hcl
	Conf.Load(
		file.NewSource(file.WithPath(GetConfigFilePath())),
		etcd.NewSource(
			etcd.WithAddress(GetEtcdAddr()...),
			etcd.StripPrefix(true),
		),
	)
	Conf.Sync()
}

func GetConfigFilePath() string {
	// if GetMode() == "dev" {
	// 	return "/Users/shixinshuiyou/go/mayo/docker/dev/conf.yaml"
	// }
	fileAddr := fmt.Sprintf("./docker/%s/conf.yaml", GetMode())

	log.Logger.Debugf("conf init file addr : %s", fileAddr)
	return fileAddr
}

func GetEtcdAddr() []string {
	addr := os.Getenv("ETCD_ADDR")
	if addr == "" {
		return []string{"127.0.0.1:2379"}
	}
	log.Logger.Debugf("conf init etcd addr : %s", addr)
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
