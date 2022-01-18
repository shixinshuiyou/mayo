package cache

import (
	"time"

	"github.com/shixinshuiyou/mayo/tool/cache/redis"
	"github.com/shixinshuiyou/mayo/tool/log"
)

type Store interface {
	Get(key string) (s string, err error)
	Save(key, data string, t time.Duration) error
	Del(key string) error
}

type Cache struct {
	Store
}

func (c *Cache) Get(key string) (s string, err error) {
	s, err = c.Store.Get(key)
	log.Logger.Debugf("cache get by key = %s, value = %s, err = %v", key, s, err)
	return
}

func (c *Cache) Save(key, data string, t time.Duration) error {
	log.Logger.Debugf("cache set key = %s,value = %s ,time = %d", key, data, t)
	return c.Store.Save(key, data, t)
}

func (c *Cache) Del(key string) error {
	log.Logger.Debugf("del key = %s", key)
	return c.Store.Del(key)
}

func NewRedisCache() *Cache {
	c := new(Cache)
	c.Store = &redis.RedisCache{}
	return c
}
