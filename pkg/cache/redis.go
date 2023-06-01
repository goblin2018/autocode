package cache

// Todo 增加本地缓存

import (
	"auto/pkg/conf"
	"auto/pkg/log"

	"github.com/redis/go-redis/v9"
)

func NewRedis(c conf.CacheConfig) *redis.Client {
	log.Infof("Start connect Redis: %s", c.Url)

	r := redis.NewClient(&redis.Options{
		Addr:     c.Url,
		DB:       0,
		Password: c.Password, // 默认没有密码
	})
	log.Infof("Redis connect success: %s", c.Url)

	return r
}
