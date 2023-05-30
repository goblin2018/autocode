package cache

// Todo 增加本地缓存

import (
	"auto/pkg/conf"

	"github.com/redis/go-redis/v9"
)

func NewRedis(c conf.CacheConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Url,
		DB:       0,
		Password: c.Password, // 默认没有密码
	})
}
