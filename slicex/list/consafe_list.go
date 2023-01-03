package list

import "sync"

var _ IArrayList[any] = (*ConSafeArrayList[any])(nil)

type ConSafeArrayList[T any] struct {
	IArrayList[T]
	lock sync.RWMutex
}

func NewConSafeArrayList[T any](cap int) *ConSafeArrayList[T] {
	return &ConSafeArrayList[T]{IArrayList: NewArrayList[T](cap)}
}
func NewConSafeArrayListOf[T any](array []T) *ConSafeArrayList[T] {
	return &ConSafeArrayList[T]{IArrayList: NewArrayListOf[T](array)}
}

func (c *ConSafeArrayList[T]) Get(index int) (T, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.IArrayList.Get(index)
}

func (c *ConSafeArrayList[T]) Set(index int, value T) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.IArrayList.Set(index, value)
}

func (c *ConSafeArrayList[T]) Append(value T) int {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.IArrayList.Append(value)
}

func (c *ConSafeArrayList[T]) QueueJump(index int, value T) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.IArrayList.QueueJump(index, value)
}

func (c *ConSafeArrayList[T]) Delete(index int) (T, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.IArrayList.Delete(index)
}

func (c *ConSafeArrayList[T]) Len() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.IArrayList.Len()
}

func (c *ConSafeArrayList[T]) Cap() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.IArrayList.Cap()
}

func (c *ConSafeArrayList[T]) Range(f func(index int, value T) error) error {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.IArrayList.Range(f)
}

func (c *ConSafeArrayList[T]) GetSlice() []T {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.IArrayList.GetSlice()
}
