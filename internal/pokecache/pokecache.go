package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cachedURLs map[string]cacheEntry
	mu         sync.RWMutex
	interval   time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cachedURLs: make(map[string]cacheEntry),
		interval:   interval,
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cachedURLs[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.cachedURLs[key]
	if !ok {
		return nil, false
	}
	return val.val, true
}

func (c *Cache) reapLoop() {
	// TODO
}
