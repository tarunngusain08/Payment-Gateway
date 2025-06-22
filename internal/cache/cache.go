package cache

import (
	"context"
	"sync"
	"time"
)

type CacheStore interface {
	Get(ctx context.Context, key string) (interface{}, bool)
	Set(ctx context.Context, key string, value interface{}) error
}

type memoryCache struct {
	data map[string]cacheItem
	mu   sync.RWMutex
	stop chan struct{}
	ttl  time.Duration
}

type cacheItem struct {
	value  interface{}
	expiry time.Time
}

func NewMemoryCacheWithJanitor(janitorInterval time.Duration, ttl time.Duration) CacheStore {
	c := &memoryCache{
		data: make(map[string]cacheItem),
		stop: make(chan struct{}),
		ttl:  ttl,
	}
	go c.startJanitor(janitorInterval)
	return c
}

func NewMemoryCacheWithTTL(janitorInterval, ttl time.Duration) CacheStore {
	c := &memoryCache{
		data: make(map[string]cacheItem),
		stop: make(chan struct{}),
		ttl:  ttl,
	}
	go c.startJanitor(janitorInterval)
	return c
}

func (c *memoryCache) Get(ctx context.Context, key string) (interface{}, bool) {
	c.mu.RLock()
	item, found := c.data[key]
	c.mu.RUnlock()
	if !found || time.Now().After(item.expiry) {
		return nil, false
	}
	return item.value, true
}

func (c *memoryCache) Set(ctx context.Context, key string, value interface{}) error {
	c.mu.Lock()
	c.data[key] = cacheItem{
		value:  value,
		expiry: time.Now().Add(c.ttl),
	}
	c.mu.Unlock()
	return nil
}

// startJanitor runs a background goroutine to remove expired items periodically.
func (c *memoryCache) startJanitor(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			now := time.Now()
			c.mu.Lock()
			for k, v := range c.data {
				if now.After(v.expiry) {
					delete(c.data, k)
				}
			}
			c.mu.Unlock()
		case <-c.stop:
			return
		}
	}
}

// Optionally, add a Close method to stop the janitor if needed.
func (c *memoryCache) Close() {
	close(c.stop)
}
