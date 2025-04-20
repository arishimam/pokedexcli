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
	mx    *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	// call cache.reapLoop()
	c := Cache{
		cache: map[string]cacheEntry{},
		mx:    &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()

	if entry, exists := c.cache[key]; exists {
		return entry.val, true
	}

	return []byte{}, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mx.Lock()
		defer c.mx.Unlock()

		for k, v := range c.cache {
			if time.Since(v.createdAt) > interval {
				delete(c.cache, k)
			}
		}

	}

}
