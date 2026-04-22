package main

import (
	"fmt"
	"sync"
	"time"
)

var _ Cache = InMemoryCache{} // это трюк для проверки типа: до тех пор пока InMemoryCache не будет реализовывать интерфейс Cache, программа не запустится

type CacheEntry struct {
	settledAt time.Time
	value     any
}

type Cache interface {
	Set(key string, value any)
	Get(key string) any
}

type InMemoryCache struct {
	mutexRW     sync.RWMutex
	m           map[string]CacheEntry
	cacheExpire time.Duration
}

func (imc InMemoryCache) Set(key string, value any) {
	var ce CacheEntry = CacheEntry{time.Now(), value}
	imc.mutexRW.Lock()
	imc.m[key] = ce
	imc.mutexRW.Unlock()
}

func (imc InMemoryCache) Get(key string) any {
	var result any

	imc.mutexRW.RLock()
	since := time.Since(imc.m[key].settledAt)
	result = imc.m[key].value
	imc.mutexRW.RUnlock()

	if since > imc.cacheExpire {
		return nil
	}
	return result

}

func NewInMemoryCache(expireIn time.Duration) *InMemoryCache {
	return &InMemoryCache{
		cacheExpire: expireIn,
		m:           make(map[string]CacheEntry),
	}
}

func main() {
	var c InMemoryCache = *NewInMemoryCache(time.Minute * 1)

	go c.Set("test", 150)
	go c.Set("test", 140)
	go c.Set("test", 130)
	go c.Set("test", 120)
	time.Sleep(1 * time.Second)
	fmt.Println(c.Get("test"))
}
