package main

import (
	"fmt"
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
	m           map[string]CacheEntry
	cacheExpire time.Duration
}

func (imc InMemoryCache) Set(key string, value any) {
	var ce CacheEntry = CacheEntry{time.Now(), value}

	imc.m[key] = ce
}

func (imc InMemoryCache) Get(key string) any {
	since := time.Since(imc.m[key].settledAt)

	if since > imc.cacheExpire {
		return nil
	}
	return imc.m[key].value

}

func NewInMemoryCache(expireIn time.Duration) *InMemoryCache {
	return &InMemoryCache{
		cacheExpire: expireIn,
		m:           make(map[string]CacheEntry),
	}
}

func main() {
	var c InMemoryCache = *NewInMemoryCache(time.Minute * 1)

	c.Set("test", 150)
	fmt.Println(c.Get("test"))
}
