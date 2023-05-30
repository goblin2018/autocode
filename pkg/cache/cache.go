package cache

import (
	"auto/pkg/conf"
	"auto/pkg/e"
	"auto/pkg/syncx"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	defaultExpiry         = time.Hour * 24 * 7
	defaultNotFoundExpiry = time.Minute
	notFoundPlaceHolder   = "*"
)

// Todo 研究 SingleFlight 在大量失败情况的表现

type Cache struct {
	barrier syncx.SingleFlight
	redis   *redis.Client
}

func NewCache(c conf.CacheConfig) *Cache {
	return &Cache{
		barrier: syncx.NewSingleFlight(),
		redis:   NewRedis(c),
	}
}

func (c *Cache) GetKey(prefix string, keys ...interface{}) string {
	var keyStr string
	for _, key := range keys {
		keyStr += fmt.Sprintf("%v:", key)
	}
	keyStr = keyStr[:len(keyStr)-1]
	return fmt.Sprintf("%s%s", prefix, keyStr)
}
func (c *Cache) GetCtx(ctx context.Context, key string, v interface{}) (err error) {
	str, err := c.redis.Get(ctx, key).Result()
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(str), v)
	if err != nil {
		return
	}
	return
}

func (c *Cache) Get(key string, v interface{}) (err error) {
	return c.GetCtx(context.Background(), key, v)
}

func (c *Cache) Set(key string, value interface{}) error {
	return c.SetCtx(context.Background(), key, value)
}

func (c *Cache) Take(
	key string, v interface{},
	query func(v interface{}) error,
	load func(v interface{}) error,
) error {
	return c.TakeCtx(context.Background(), key, v, query, load)
}

func (c *Cache) TakeCtx(
	ctx context.Context, key string, v interface{},
	query func(v interface{}) error,
	load func(v interface{}) error,
) error {
	_, err := c.barrier.Do(key, func() (interface{}, error) {
		err := c.GetCtx(ctx, key, v)
		if err != nil {
			if err == e.NotExists {
				return nil, e.NotExists
			} else if err != redis.Nil {
				return nil, err
			}
		}
		err = query(v)
		if err != nil {
			return nil, err
		}
		err = load(v)
		return nil, err
	})

	return err
}

func (c *Cache) SetCtx(ctx context.Context, key string, value interface{}) error {
	d, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.redis.Set(ctx, key, string(d), unstableDuration(defaultExpiry)).Err()
}

func (c *Cache) Del(key string) error {
	return c.DelCtx(context.Background(), key)
}

func (c *Cache) DelCtx(ctx context.Context, key string) error {
	return c.redis.Del(ctx, key).Err()
}

func (c *Cache) Incr(key string) error {
	return c.redis.Incr(context.Background(), key).Err()
}

func (c *Cache) IncrCtx(ctx context.Context, key string) error {
	return c.redis.Incr(ctx, key).Err()
}

func (c *Cache) Decr(key string) error {
	return c.redis.Decr(context.Background(), key).Err()
}

func (c *Cache) DecrCtx(ctx context.Context, key string) error {
	return c.redis.Decr(ctx, key).Err()
}

func (c *Cache) IncrBy(key string, value int64) error {
	return c.redis.IncrBy(context.Background(), key, value).Err()
}

func (c *Cache) IncrByCtx(ctx context.Context, key string, value int64) error {
	return c.redis.IncrBy(ctx, key, value).Err()
}
