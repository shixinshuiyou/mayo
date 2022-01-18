package redis

import (
	"time"

	"github.com/go-redis/redis"
)

type RedisCache struct {
	Client *redis.Client
}

func (rc *RedisCache) Get(k string) (s string, err error) {
	return getInstance().Get(k).Result()
}

// TODO d interface{} 可以做类型断言-实现set zet 等
func (rc *RedisCache) Save(k, d string, t time.Duration) error {
	return getInstance().Set(k, d, t).Err()
}

func (rc *RedisCache) Del(k string) error {
	return getInstance().Del(k).Err()
}
