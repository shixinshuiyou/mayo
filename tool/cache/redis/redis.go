package redis

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/shixinshuiyou/mayo/tool/config"
	"github.com/shixinshuiyou/mayo/tool/log"
)

var (
	single *redis.Client
	once   sync.Once
	option *redis.Options
)

func init() {
	host := config.Conf.Get("redis-cli", "host").String("127.0.0.1")
	port := config.Conf.Get("redis-cli", "port").Int(6379)
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Logger.Debugf("redis session connt addr :%s", addr)
	option = &redis.Options{
		Addr:       addr,
		Password:   config.Conf.Get("redis", "password").String(""),
		DB:         0,
		MaxConnAge: time.Duration(config.Conf.Get("redis", "maxconnage").Int(10)) * time.Second,
	}
}

func getInstance() *redis.Client {
	once.Do(func() {
		single = redis.NewClient(option)
	})

	return single
}

func GetInstance() *redis.Client {
	return getInstance()
}
