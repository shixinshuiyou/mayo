package cache

import (
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
	addr := config.Conf.Get("redis", "address").String("127.0.0.1:6379")
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

// TODO
func InitRedisConf(op *redis.Options) {
	option = op
}
