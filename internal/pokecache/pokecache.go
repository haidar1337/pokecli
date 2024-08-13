package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	if c.entries == nil {
		c.entries = make(map[string]cacheEntry)
	}
	c.entries[key] = cacheEntry{
		val: val,
		createdAt: time.Now(),
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry, ok := c.entries[key]
	c.mu.Unlock()
	return entry.val, ok
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{}
	go c.reapLoop(interval)
	return &c
}


func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			c.mu.Lock()
			fmt.Printf("Cache check, t value: %v", t)
			fmt.Println()
			for key, val := range c.entries {
				if val.createdAt.Sub(t) <= 0 {
					fmt.Printf("Deleted %s from cache after %v", key, val.createdAt.Sub(t))
					delete(c.entries, key)
				}
			}
			c.mu.Unlock()
		}
	}
}