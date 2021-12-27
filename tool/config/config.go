package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/etcd"
	"github.com/micro/go-micro/v2/config/source/file"
	"github.com/shixinshuiyou/mayo/tool/log"
)

// 采用etcd-conf 作为etcd 配置中心
var (
	Conf config.Config
)

func init() {
	Conf, _ = config.NewConfig()
	confPath := GetConfigFilePath()
	log.Logger.Debugf("load file config,file path : %s", confPath)
	// 加载文件配置
	fileSource := file.NewSource(file.WithPath(confPath))
	// 加载和合并多个源。合并优先级顺序相反. 此处  etcd > file
	// 官方支持的解析器还有 yaml、toml、xml、hcl
	Conf.Load(fileSource)

	etcdAddr := Conf.Get("etcd", "address").String("127.0.0.1:2380")
	log.Logger.Infof("load etcd config,etcd addr : %s", etcdAddr)
	etcdAddrs := strings.Split(etcdAddr, ",")

	etcdSource := etcd.NewSource(
		etcd.WithAddress(etcdAddrs...),
		etcd.WithPrefix("/micro/config"),
		etcd.StripPrefix(true),
	)

	Conf.Load(etcdSource)

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
