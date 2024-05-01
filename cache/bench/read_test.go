// При обычном (а не параллельном) чтении лучше использовать Mutex.
package cache_benchs

import (
	"context"
	"testing"

	"structures/cache"
	"structures/cache/cache_mutex"
	"structures/cache/cache_rwmutex"
)

// 87608             13721 ns/op               0 B/op         0 allocs/op
func Benchmark_Mutex_Read(t *testing.B) {
	c := cache_mutex.NewCache[int, int]()
	fillCache(c)
	t.ResetTimer()
	for range t.N {
		read(c)
	}
}

// 85665             13941 ns/op               0 B/op         0 allocs/op
func Benchmark_RWMutex_Read(t *testing.B) {
	c := cache_rwmutex.NewCache[int, int]()
	fillCache(c)
	t.ResetTimer()
	for range t.N {
		read(c)
	}
}

// 2239            536290 ns/op          152011 B/op       3000 allocs/op
func Benchmark_Mutex_Read_WithCtx(t *testing.B) {
	c := cache_mutex.NewCache[int, int]()
	fillCache(c)
	t.ResetTimer()
	for range t.N {
		read_WithContext(c)
	}
}

// 2223            540658 ns/op          152023 B/op       3000 allocs/op
func Benchmark_RWMutex_Read_WithCtx(t *testing.B) {
	c := cache_rwmutex.NewCache[int, int]()
	fillCache(c)
	t.ResetTimer()
	for range t.N {
		read_WithContext(c)
	}
}

func read(c cache.Cache[int, int]) {
	for i := 0; i < 1000; i++ {
		if _, ok := c.Get(i); !ok {
			panic("Ключ не найден")
		}
	}
}

func read_WithContext(c cache.Cache[int, int]) {
	for i := 0; i < 1000; i++ {
		vPtr, _ := cache.CacheGetWithContext(context.Background(), i, c)
		if vPtr == nil {
			panic("Ключ не найден")
		}
	}
}
