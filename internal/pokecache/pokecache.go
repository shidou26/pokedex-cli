package pokecache

import (
	"time"
	"sync"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]CacheEntry
	mutex *sync.Mutex
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mutex.Lock()
		removedKey := []string{}
		for key, value := range c.cache {
			if time.Since(value.createdAt) > interval {
				removedKey = append(removedKey, key)
			}
		}
		for _, key := range removedKey {
			delete(c.cache, key)
		}
		c.mutex.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache {
		cache: make(map[string]CacheEntry),
		mutex: &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[key] = CacheEntry{time.Now(), val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	val, ok := c.cache[key]
	return val.val, ok
}
