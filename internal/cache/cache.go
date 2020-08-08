package cache

import "time"

type Cache interface {
	Ping() error
	Get(key string) (string, error)
	Set(key string, value interface{}, ttl time.Duration) error
	// Invalidate invalidates key by given names. Pattern shouldn't contains redis prefix!
	Invalidate(pattern string) error
}
