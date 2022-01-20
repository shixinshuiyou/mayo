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

type redisConf struct {
	Host     string
	Port     int
	Password string
}

func init() {
	conf := new(redisConf)
	config.Conf.Get("redis-cli").Scan(conf)

	addr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	log.Logger.Debugf("redis session connt addr :%s-%s,", addr, conf.Password)
	option = &redis.Options{
		Addr:       addr,
		Password:   conf.Password,
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
