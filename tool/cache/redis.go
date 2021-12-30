package cache

import (
	"sync"

	"github.com/go-redis/redis"
)

var (
	single *redis.Client
	once   sync.Once
	option *redis.Options
)

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
