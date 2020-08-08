package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type RedisCache struct {
	client *redis.Client
	prefix string
}

func NewRedisCache(client *redis.Client, prefix string) *RedisCache {
	return &RedisCache{
		client: client,
		prefix: prefix,
	}
}

// Ping method check connection to redis.
func (r RedisCache) Ping() error {
	if r.client == nil {
		return errors.New("failed ping redis. Redis client isn't initialized")
	}
	_, err := r.client.Ping().Result()
	return err
}

func (r RedisCache) Get(key string) (string, error) {
	if r.client == nil {
		return "", fmt.Errorf("failed get key %s from redis: client isn't initialized",
			key)
	}
	key = r.resolvePrefix(key)

	return r.client.Get(key).Result()
}

func (r RedisCache) Set(key string, value interface{}, ttl time.Duration) error {
	if r.client == nil {
		return fmt.Errorf("failed set key %s in redis: client isn't initialized",
			key)
	}
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}

	key = r.resolvePrefix(key)

	return r.client.Set(key, b, ttl).Err()
}

func (r RedisCache) Invalidate(pattern string) error {
	if r.client == nil {
		return fmt.Errorf("failed invalidate keys by pattern"+
			"%s in redis: client isn't initialized",
			pattern)
	}
	p := r.resolvePrefix(pattern)
	result := r.client.Keys(p)
	if result.Err() != nil {
		return fmt.Errorf("failed receive keys for invalidate. Error: %s", result.Err())
	}
	keys := result.Val()
	if len(keys) == 0 {
		return nil
	}
	return r.client.Del(keys...).Err()
}

func (r RedisCache) resolvePrefix(key string) string {
	if len(key) > 0 {
		return fmt.Sprintf("%s:%s", r.prefix, key)
	}
	return key
}
