package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]CacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cacheMap: map[string]CacheEntry{},
		interval: interval,
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.cacheMap[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.cacheMap[key]
	if ok {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		defer c.mu.Unlock()
		for k, v := range c.cacheMap {
			if time.Since(v.createdAt) > c.interval {
				delete(c.cacheMap, k)
			}
		}
	}
}
