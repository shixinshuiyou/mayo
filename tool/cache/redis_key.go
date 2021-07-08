package cache

import "fmt"

type Cache struct {
	KeyExpression string // 构成key的表达式
}

func (c *Cache) GetRedisKey(val ...string) string {
	return fmt.Sprintf(c.KeyExpression, val)
}

func (c *Cache) SetRedisPrefix() {

}
