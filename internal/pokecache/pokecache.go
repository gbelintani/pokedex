package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cache: map[string]cacheEntry{},
		mu:    &sync.RWMutex{},
	}

	c := time.Tick(interval)
	go func() {
		for next := range c {
			cache.reapLoop(next, interval)
		}
	}()

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, exists := c.cache[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(now time.Time, interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, value := range c.cache {
		if now.Sub(value.createdAt) >= interval {
			delete(c.cache, key)
			fmt.Printf("removing cache entry %s\n", key)
		}
	}
}
