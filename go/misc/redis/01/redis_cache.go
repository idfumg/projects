package main

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v9"
)

type RedisCache struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisCache(client *redis.Client, ttl time.Duration) *RedisCache {
	return &RedisCache{
		client: client,
		ttl:    ttl,
	}
}

func (c *RedisCache) Get(key int) (string, bool) {
	ctx := context.Background()
	k := strconv.Itoa(key)
	val, err := c.client.Get(ctx, k).Result()
	return val, err == nil
}

func (c *RedisCache) Set(key int, val string) error {
	ctx := context.Background()
	k := strconv.Itoa(key)
	_, err := c.client.Set(ctx, k, val, c.ttl).Result()
	return err
}

func (c *RedisCache) Remove(key int) error {
	ctx := context.Background()
	k := strconv.Itoa(key)
	_, err := c.client.Del(ctx, k).Result()
	return err
}
