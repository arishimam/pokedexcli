package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mx    sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	// call cache.reapLoop()
	newCache := &Cache{cache: map[string]cacheEntry{}}
	newCache.reapLoop(interval)
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mx.Lock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mx.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mx.Lock()
	if entry, ok := c.cache[key]; ok {
		c.mx.Unlock()
		return entry.val, true
	}

	c.mx.Unlock()
	return []byte{}, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	c.mx.Lock()
	defer c.mx.Unlock()

	for k, v := range c.cache {
		if time.Since(v.createdAt) > interval {
			delete(c.cache, k)
		}
	}
}
