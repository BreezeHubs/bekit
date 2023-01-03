package list

import "github.com/BreezeHubs/bekit/slicex"

var _ IArrayList[any] = (*LinkedList[any])(nil)

//双向链表节点数据
type node[T any] struct {
	prev  *node[T]
	next  *node[T]
	value T
}

// LinkedList 双向循环链表
type LinkedList[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

func NewLinkedList[T any]() IArrayList[T] {
	head := &node[T]{}
	tail := &node[T]{next: head, prev: head}
	head.next, head.prev = tail, tail
	return &LinkedList[T]{
		head: head,
		tail: tail,
	}
}

func (l *LinkedList[T]) isOutOfRange(index int) error {
	if index < 0 || index >= l.length {
		return slicex.IndexError
	}
	return nil
}

func (l *LinkedList[T]) findNode(index int) *node[T] {
	//使用 二分法 查找
	var tmp *node[T]
	if index <= l.Len()/2 {
		tmp = l.head
		for i := 0; i <= index; i++ {
			tmp = tmp.next
		}
	} else {
		tmp = l.tail
		for i := 0; i <= l.Len()-index; i++ {
			tmp = tmp.prev
		}
	}
	return tmp
}

func (l *LinkedList[T]) Get(index int) (value T, err error) {
	if err = l.isOutOfRange(index); err != nil {
		return
	}
	value = l.findNode(index).value
	return
}

func (l *LinkedList[T]) Set(index int, value T) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[T]) Append(value T) int {
	node := &node[T]{prev: l.tail.prev, next: l.tail, value: value}
	node.prev.next, node.next.prev = node, node
	l.length++
	return 0
}

func (l *LinkedList[T]) QueueJump(index int, value T) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[T]) Delete(index int) (T, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[T]) Len() int {
	return l.length
}

func (l *LinkedList[T]) Cap() int {
	return l.Len()
}

func (l *LinkedList[T]) Range(f func(index int, value T) error) error {
	tmp := l.head
	for i := 0; i < l.Len(); i++ {
		if err := f(i, tmp.value); err != nil {
			return err
		}
		tmp = tmp.next
	}
	return nil
}

func (l *LinkedList[T]) GetSlice() []T {
	//TODO implement me
	panic("implement me")
}
