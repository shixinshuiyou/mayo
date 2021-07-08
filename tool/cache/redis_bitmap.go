package cache

import (
	"time"

	"github.com/go-redis/redis"
)

type RedisBitMap struct {
	RedisKey string
	Ttl      time.Duration
}

func NewRedisBitMap(key string, ttl time.Duration) *RedisBitMap {
	return &RedisBitMap{RedisKey: key, Ttl: ttl}
}

func (r *RedisBitMap) GetBit(offset int) int {
	val := RedisInstance().GetBit(r.RedisKey, int64(offset)).Val()
	return int(val)
}

func (r *RedisBitMap) SetBit(offset, val int) {
	redisIns := RedisInstance()
	redisIns.SetBit(r.RedisKey, int64(offset), val)
	if r.Ttl > 0 {
		redisIns.Expire(r.RedisKey, r.Ttl)
	}
}

func (r *RedisBitMap) BitCount(start, end int) int {
	val := RedisInstance().BitCount(r.RedisKey, &redis.BitCount{Start: int64(start), End: int64(end)}).Val()
	return int(val)
}
