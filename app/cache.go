package app

import "github.com/garyburd/redigo/redis"

type Cache interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Del(key string) error
	Close()
}

func NewCache(client *redis.Pool) Cache {
	return &CacheImpl{client}
}

type CacheImpl struct {
	client *redis.Pool
}

func (c *CacheImpl) Get(key string) (string, error) {
	return redis.String(c.client.Get().Do("GET", key))
}

func (c *CacheImpl) Set(key string, value string) error {
	_, err := c.client.Get().Do("SET", key, value)
	return err
}

func (c *CacheImpl) Del(key string) error {
	_, err := c.client.Get().Do("DEL", key)
	return err
}

func (c *CacheImpl) Close() {
	c.client.Close()
}
