package list

import "github.com/BreezeHubs/bekit/slicex"

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
	//valuesTmp := a.values[index+1:]
	//a.values = append(a.values[:index], value)
	//a.values = append(a.values, valuesTmp...)
	panic("implement me")
}

func (a *ArrayList[T]) Delete(index int) error {
	//TODO implement me
	panic("implement me")
}

func (a *ArrayList[T]) Len() int {
	//TODO implement me
	panic("implement me")
}

func (a *ArrayList[T]) Cap() int {
	//TODO implement me
	panic("implement me")
}

func (a *ArrayList[T]) Range(f func(index int, value T) error) error {
	//TODO implement me
	panic("implement me")
}

func NewArrayList[T any](cap int) IArrayList[T] {
	return &ArrayList[T]{values: make([]T, 0, cap)}
}
