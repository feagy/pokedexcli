package main

import (
    "sync"
    "time"
)
type cacheEntry struct {
    createdAt time.Time
    val []byte
}

type cache struct {
    cacheEntries map[string]cacheEntry
    interval time.Duration
    mu *sync.Mutex
}

func newCache(interval time.Duration) *cache {
    c := &cache{
        cacheEntries: map[string]cacheEntry{},
        interval: interval,
        mu: &sync.Mutex{},
    }   

    go c.reapLoop()
    return c
}

func (c *cache) Add(key string, val []byte) {
    c.mu.Lock()
    c.cacheEntries[key] = cacheEntry{
        createdAt: time.Now(),
        val: val,
    }
    c.mu.Unlock()
}

func (c *cache) Get(key string) ([]byte, bool) {
    c.mu.Lock()
    cEntry, ok := c.cacheEntries[key]
    c.mu.Unlock()
    return cEntry.val, ok
}

func (c *cache) reapLoop() {
    ticker := time.NewTicker(c.interval)
    for range ticker.C {
        c.mu.Lock()  
        for k, v := range c.cacheEntries {
            if time.Since(v.createdAt) > c.interval {
                delete(c.cacheEntries, k)
            }
        }
        c.mu.Unlock()
    }
}
