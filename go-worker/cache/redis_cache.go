package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(c *redis.Client) *RedisCache {
	return &RedisCache{
		client: c,
	}
}

func (c *RedisCache) Get(ctx context.Context, key string, target interface{}) bool {
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return false
	}
	if err := json.Unmarshal([]byte(val), target); err != nil {
		return false
	}
	return true
}

func (c *RedisCache) Set(ctx context.Context, key string, val interface{}, ttl time.Duration) error {
	stringifyVal, err := json.Marshal(val)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, key, stringifyVal, ttl).Err()
}

func (c *RedisCache) HasSet(ctx context.Context, setName string, val string) bool {
	if exists, _ := c.client.SIsMember(ctx, setName, val).Result(); exists {
		return true
	}
	return false // Todo maybe check if err and return accordingly
}

func (c *RedisCache) AddSet(ctx context.Context, setName string, val string) error {
	return c.client.SAdd(ctx, setName, val).Err()
}

func (c *RedisCache) RemoveSet(ctx context.Context, setName string, val string) error {
	return c.client.SRem(ctx, setName, val).Err()
}
