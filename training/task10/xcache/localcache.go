package xcache

import (
	"fmt"
	"gitlab.myteksi.net/gophers/go/commons/util/cache/lrucache"
	"time"

	"gitlab.myteksi.net/gophers/go/commons/util/cache/localcache"
)

func cacheLoader(key string) (localcache.Value, error) {
	// The value could be loaded from the db here, and the localcache
	// package ensures that the function will only be called once
	return "value", nil
}

func TestLocalCache() {
	cacheBuilder := localcache.CacheBuilder{}
	cache := cacheBuilder.
		ExpireAfterAccess(10 * time.Millisecond).
		ExpireAfterWrite(10 * time.Millisecond).
		RecordStats("cache").
		SetCacheLoader(cacheLoader).
		Build()

	v, hit, err := cache.Get("key")
	if err != nil {
		panic(err)
	}
	if hit {
		fmt.Println("Loaded:", v)
	}
}

func TestLRUCache() {
	cache, _ := lrucache.New(1)
	cache.Insert("key1", "value1")
	cache.Insert("key2", "value2")
	val, _ := cache.Get("key1")
	fmt.Println("Value 1:", val)

	val2, _ := cache.Get("key2")
	fmt.Println("Value 2:", val2)
}