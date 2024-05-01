// 1. Операций +- одинаково -> Mutex
package cache_benchs

import (
	"math/rand"
	"structures/cache"
	"structures/cache/cache_mutex"
	"structures/cache/cache_rwmutex"
	"sync"
	"testing"
)

// 1489            819628 ns/op             520 B/op          9 allocs/op
func Benchmark_Mutex_SetGetDelete(t *testing.B) {
	c := cache_mutex.NewCache[int, int]()
	t.ResetTimer()
	for range t.N {
		getSetDeleteOperations(c)
	}
}

// 1426            826481 ns/op             515 B/op          9 allocs/op
func Benchmark_RWMutex_SetGetDelete(t *testing.B) {
	c := cache_rwmutex.NewCache[int, int]()
	t.ResetTimer()
	for range t.N {
		getSetDeleteOperations(c)
	}
}

// func Benchmark_Atomic_SetGetDelete(t *testing.B) {
// 	for range t.N {

// 	}
// }

func getSetDeleteOperations(c cache.Cache[int, int]) {
	N := 1000
	wg := new(sync.WaitGroup)
	for range 3 {
		wg.Add(1)
		go func() {
			for i := range N {
				c.Set(i, i)
			}
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			for i := range N {
				c.Get(i)
			}
			wg.Done()
		}()

	}
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-done:
				wg.Done()
				return
			default:
				key := rand.Intn(N)
				if _, ok := c.Get(key); ok {
					c.Delete(key)
				}
			}
		}
	}()
	wg.Wait()
	wg.Add(1)
	done <- struct{}{}
	wg.Wait()
}
