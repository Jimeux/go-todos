package common

import (
	"github.com/go-redis/redis"
	"time"
)

const (
	DefaultTimeout = 100 * time.Second
)

type Cache interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Del(key string) error
	Close()
}

func NewCache(client *redis.Client) Cache {
	return &CacheImpl{client}
}

type CacheImpl struct {
	client *redis.Client
}

func (c *CacheImpl) Get(key string) (string, error) {
	return c.client.Get(key).Result()
}

func (c *CacheImpl) Set(key string, value string) error {
	return c.client.Set(key, value, DefaultTimeout).Err()
}

func (c *CacheImpl) Del(key string) error {
	return c.client.Del(key).Err()
}

func (c *CacheImpl) Close() {
	c.client.Close()
}
