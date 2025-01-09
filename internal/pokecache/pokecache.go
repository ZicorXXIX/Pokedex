package pokecache

import (
	"sync"
	"time"
    "fmt"
)


type Cache struct {
    mu          sync.Mutex
    entries     map[string]cacheEntry
    interval    time.Duration
}

type cacheEntry struct {
    createdAt   time.Time
    val         []byte
}


func NewCache(interval time.Duration) *Cache {
    
    c :=Cache{
        entries: map[string]cacheEntry{},
        interval: interval,
    }
    go c.reapLoop()
    return &c
}

func (c *Cache) Add(key string, value []byte) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.entries[key] = cacheEntry{
        createdAt: time.Now(),
        val: value,
    }
}

func (c *Cache) Get(key string) ([]byte, bool){
    c.mu.Lock()
    defer c.mu.Unlock()
    cacheValue, exists :=c.entries[key]
    if !exists {
        return nil, exists
    }
    return cacheValue.val, exists
}

func (c *Cache) reapLoop() {
    ticker := time.NewTicker(c.interval) 
    defer ticker.Stop()
    
    for{
        select {
        case <-ticker.C:
            fmt.Println("Clearing Cache")
            c.mu.Lock()
            c.mu.Unlock()
            for key, entry := range c.entries {
                if time.Since(entry.createdAt) > c.interval{
                    delete(c.entries, key)
                }
            }
        default:
            //do nothing
        }
    }
}
