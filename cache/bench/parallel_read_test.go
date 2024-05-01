// 1. Больше Get, чем Set -> RWMutex в 2 раза быстрее.
// 2. Get_Ctx -> нет явного победителя.
package cache_benchs

import (
	"context"
	"runtime"
	"sync"
	"testing"

	"structures/cache"
	"structures/cache/cache_mutex"
	"structures/cache/cache_rwmutex"
)

// 4665            257056 ns/op             524 B/op         15 allocs/op
func Benchmark_Mutex_Read_Parallel(t *testing.B) {
	c := cache_mutex.NewCache[int, int]()
	fillCache(c)
	t.ResetTimer()
	for range t.N {
		parallelRead(c)
	}
}

// 8428            136268 ns/op             521 B/op         15 allocs/op
func Benchmark_RWMutex_Read_Parallel(t *testing.B) {
	c := cache_rwmutex.NewCache[int, int]()
	fillCache(c)
	t.ResetTimer()
	for range t.N {
		parallelRead(c)
	}
}

// 2224            539750 ns/op         532539 B/op       10515 allocs/op
func Benchmark_Mutex_Read_WithCtx_Parallel(t *testing.B) {
	c := cache_mutex.NewCache[int, int]()
	fillCache(c)
	t.ResetTimer()
	for range t.N {
		parallelRead_WithContext(c)
	}
}

// 2280            526539 ns/op         532539 B/op       10515 allocs/op
func Benchmark_RWMutex_Read_WithCtx_Parallel(t *testing.B) {
	c := cache_rwmutex.NewCache[int, int]()
	fillCache(c)
	t.ResetTimer()
	for range t.N {
		parallelRead_WithContext(c)
	}
}

func fillCache(c cache.Cache[int, int]) {
	for i := 0; i < 1000; i++ {
		c.Set(i, i)
	}
}

func parallelRead(c cache.Cache[int, int]) {
	wg := new(sync.WaitGroup)
	for i := range runtime.NumCPU() - 1 {
		offset := i * 10
		wg.Add(1)
		go func(j int) {
			for i := offset; i < offset+500; i++ {
				if _, ok := c.Get(i); !ok {
					panic("Ключ не найден")
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func parallelRead_WithContext(c cache.Cache[int, int]) {
	wg := new(sync.WaitGroup)
	for i := range runtime.NumCPU() - 1 {
		offset := i * 10
		wg.Add(1)
		go func(j int) {
			for i := offset; i < offset+500; i++ {
				vPtr, _ := cache.CacheGetWithContext(context.Background(), i, c)
				if vPtr == nil {
					panic("Ключ не найден")
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
