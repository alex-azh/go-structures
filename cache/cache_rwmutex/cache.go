package cache_rwmutex

import "sync"

// Структура для хранения ключ-значения с конкурентным доступом.
type Cache[TSource comparable, TResult any] struct {
	mu *sync.RWMutex
	m  map[TSource]TResult
}

// Конструктор для создания нового кеша.
func NewCache[TSource comparable, TResult any]() *Cache[TSource, TResult] {
	return &Cache[TSource, TResult]{
		m:  make(map[TSource]TResult, 0),
		mu: new(sync.RWMutex),
	}
}

// Получить значение по ключу.
func (c *Cache[TSource, TResult]) Get(key TSource) (TResult, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.m[key]
	return v, ok
}

// Задать значение по ключу.
func (c *Cache[TSource, TResult]) Set(key TSource, value TResult) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = value
}

// Удалить значение по ключу.
func (c *Cache[TSource, TResult]) Delete(key TSource) {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.m[key]
	if ok {
		delete(c.m, key)
	}
}
