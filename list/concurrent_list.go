package list

import "sync"

var (
	_ List[any] = &ConcurrentList[any]{}
)

type ConcurrentList[T any] struct {
	List[T]
	lock sync.RWMutex
}

func (c *ConcurrentList[T]) Get(idx int) (T, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.List.Get(idx)
}

func (c *ConcurrentList[T]) Append(ts ...T) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	return c.List.Append(ts...)
}

func (c *ConcurrentList[T]) Add(idx int, t T) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	return c.List.Add(idx, t)
}

func (c *ConcurrentList[T]) Set(idx int, t T) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	return c.List.Set(idx, t)
}

func (c *ConcurrentList[T]) Delete(idx int) (T, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	return c.List.Delete(idx)
}

func (c *ConcurrentList[T]) Len() int {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.List.Len()
}

func (c *ConcurrentList[T]) Cap() int {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.List.Cap()
}

func (c *ConcurrentList[T]) Range(fn func(idx int, t T) error) error {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.List.Range(fn)
}

func (c *ConcurrentList[T]) AsSlice() []T {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.List.AsSlice()
}
