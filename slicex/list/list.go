package list

import (
	"github.com/BreezeHubs/bekit/slicex"
)

var _ IArrayList[any] = (*ArrayList[any])(nil)

type ArrayList[T any] struct {
	values []T
}

func (a *ArrayList[T]) isOutOfRange(index int) error {
	if index < 0 || index >= len(a.values) {
		return slicex.IndexError
	}
	return nil
}

func (a *ArrayList[T]) Get(index int) (value T, err error) {
	if err = a.isOutOfRange(index); err != nil {
		return
	}
	value = a.values[index]
	return
}

func (a *ArrayList[T]) Set(index int, value T) error {
	if err := a.isOutOfRange(index); err != nil {
		return err
	}
	a.values[index] = value
	return nil
}

func (a *ArrayList[T]) Append(value T) int {
	a.values = append(a.values, value)
	return len(a.values) - 1
}

func (a *ArrayList[T]) QueueJump(index int, value T) error {
	if err := a.isOutOfRange(index); err != nil {
		return err
	}

	//新值加入数组
	a.values = append(a.values, value)
	//将index后的值往后移动一个下标
	copy(a.values[index+1:], a.values[index:])
	//将index位赋值为新值
	a.values[index] = value
	return nil
}

func (a *ArrayList[T]) Delete(index int) (delValue T, err error) {
	if err = a.isOutOfRange(index); err != nil {
		return
	}
	delValue = a.values[index]

	j := 0
	for i, v := range a.values {
		if i != index {
			a.values[j] = v
			j++
		}
	}
	a.values = a.values[:j]
	a.shrink() //控制缩容
	return
}

func (a *ArrayList[T]) shrink() {
	a.values = slicex.ArrayShrink[T](a.values)
}

func (a *ArrayList[T]) Len() int {
	return len(a.values)
}

func (a *ArrayList[T]) Cap() int {
	return cap(a.values)
}

func (a *ArrayList[T]) Range(f func(index int, value T) error) error {
	for i, t := range a.values {
		if err := f(i, t); err != nil {
			return err
		}
	}
	return nil
}

func (a ArrayList[T]) GetSlice() []T {
	return a.values
}

func NewArrayList[T any](cap int) IArrayList[T] {
	return &ArrayList[T]{values: make([]T, 0, cap)}
}

func NewArrayListOf[T any](array []T) IArrayList[T] {
	return &ArrayList[T]{values: array}
}
