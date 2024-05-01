// 1. Set -> Mutex (rwmutex хранит две очереди, по сути)

package cache_benchs

import (
	"math/rand"
	"runtime"
	"sync"
	"testing"

	"structures/cache"
	"structures/cache/cache_mutex"
	"structures/cache/cache_rwmutex"
)

// 1986            593927 ns/op            3076 B/op          8 allocs/op
func Benchmark_Mutex_Set_Parallel(t *testing.B) {
	c := cache_mutex.NewCache[int, int]()
	t.ResetTimer()
	for range t.N {
		parallelSet(c)
	}
}

// 1688            713422 ns/op            3591 B/op          8 allocs/op
func Benchmark_RWMutex_Set_Parallel(t *testing.B) {
	c := cache_rwmutex.NewCache[int, int]()
	t.ResetTimer()
	for range t.N {
		parallelSet(c)
	}
}

// func Benchmark_Atomic_Set(t *testing.B) {
// 	for range t.N {
// 	}
// }

func parallelSet(c cache.Cache[int, int]) {
	wg := new(sync.WaitGroup)
	for range runtime.NumCPU() - 3 {
		wg.Add(1)
		go func() {
			for range 1000 {
				key := rand.Intn(100_000)
				c.Set(key, 10)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
