package cache

import (
	"time"

	"github.com/shixinshuiyou/mayo/tool/config"
)

func SetMysqlCache(address string) {
	redisKey := "mysql-cli"
	expire := 5 * time.Second
	userCache.Save(redisKey, address, expire)
}

func GetMysqlCache() string {
	redisKey := "mysql-cli"
	host, _ := userCache.Get(redisKey)
	if host == "" {
		host = config.Conf.Get("dbmayo", "host").String("example")
		SetMysqlCache(host)
	}
	return host
}
