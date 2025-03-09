package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
    mux *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entries:  make(map[string]cacheEntry),
        mux: &sync.RWMutex{},
	}
	go cache.readLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
    c.mux.Lock()
    defer c.mux.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
    c.mux.RLock()
    defer c.mux.RUnlock()

	entry, exists := c.entries[key]
	if exists {
		return entry.val, true
	}
	return []byte{}, false
}

func (c *Cache) readLoop(interval time.Duration) {
    ticker := time.NewTicker(interval)
    for range ticker.C {
        c.reap(time.Now().UTC(), interval)
    }
}

func (c *Cache) reap(now time.Time, last time.Duration) {
    c.mux.Lock()
    defer c.mux.Unlock()

    for k, v := range c.entries {
        if v.createdAt.Before(now.Add(-last)) {
            delete(c.entries, k)
        }
    }
}
