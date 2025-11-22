package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu      *sync.Mutex
	entries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		mu:      &sync.Mutex{},
		entries: make(map[string]cacheEntry),
	}

	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, entry []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       entry,
	}
}

func (c *Cache) Get(key string) (cacheEntry, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.entries[key]
	if !ok {
		return cacheEntry{}, ok
	}
	return entry, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(invernal time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	toDelete := []string{}
	for k, v := range c.entries {
		if v.createdAt.Add(invernal).After(time.Now()) {
			toDelete = append(toDelete, k)
		}
	}

	for _, key := range toDelete {
		delete(c.entries, key)
	}
}
