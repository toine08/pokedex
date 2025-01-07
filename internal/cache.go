package utils

import (
	"sync"
	"time"
)

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type Cache struct {
	Mu    *sync.Mutex
	Cache map[string]CacheEntry
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{ // Assuming you have a map
		Mu:    &sync.Mutex{},               // Proper initialization of mutex as necessary
		Cache: make(map[string]CacheEntry), // Initialize the map
	}
	go cache.reapLoop(interval) // Start the reapLoop in a new goroutine
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	value := CacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
	c.Cache[key] = value
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	if cacheEntry, exists := c.Cache[key]; exists {
		return cacheEntry.Val, true
	} else {
		return nil, false
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.Mu.Lock()
		for key, entry := range c.Cache {
			if entry.CreatedAt.Add(interval).Before(time.Now()) {
				delete(c.Cache, key)
			}
		}
		c.Mu.Unlock()
	}
}
