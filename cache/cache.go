package cache

import (
	"context"
	"time"
)

type Cache[TSource comparable, TResult any] interface {
	Get(key TSource) (TResult, bool)
	Set(key TSource, value TResult)
	Delete(key TSource)
}

func CacheGetWithContext[TSource comparable, TResult any](
	ctx context.Context,
	key TSource,
	cache Cache[TSource, TResult]) (*TResult, error) {
	ch := make(chan TResult)
	go func() {
		defer close(ch)
		if v, ok := cache.Get(key); ok {
			ch <- v
		}
	}()
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case v, ok := <-ch:
		timeBefore := time.Now()
		deadLine, _ := ctx.Deadline()

		if !deadLine.IsZero() && (deadLine.UnixNano()-timeBefore.UnixNano()) < 0 {
			return nil, ctx.Err()
		}

		if ok {
			return &v, nil
		} else {
			return nil, nil
		}
	}
}
