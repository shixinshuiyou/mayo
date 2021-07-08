package cache

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

var cache *redis.Client
var once sync.Once

func getInstance() *redis.Client {
	once.Do(func() {
		host := "101.133.169.39"
		port := 6379
		password := ""
		addr := fmt.Sprintf("%s:%d", host, port)
		client := redis.NewClient(&redis.Options{
			Addr:       addr,
			Password:   password,
			DB:         0,
			MaxConnAge: 20 * time.Second,
		})
		cache = client
	})
	return cache
}

func RedisInstance() *redis.Client {
	return getInstance()
}
