package list

import (
	"github.com/BreezeHubs/bekit/slicex"
)

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

//判断是否超出下标
func (l *LinkedList[T]) isOutOfRange(index int) error {
	if index < 0 || index >= l.length {
		return slicex.IndexError
	}
	return nil
}

//根据下标找到对应节点
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
		for i := l.Len(); i > index; i-- {
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

func (l *LinkedList[T]) Set(index int, value T) (err error) {
	if err = l.isOutOfRange(index); err != nil {
		return
	}
	l.findNode(index).value = value
	return
}

func (l *LinkedList[T]) Append(value T) int {
	node := &node[T]{
		prev:  l.tail.prev, //为尾节点的上一个
		next:  l.tail,      //为尾节点
		value: value,
	}
	node.prev.next = node //将上一个节点的下一个更新为新节点的指针
	node.next.prev = node //将下一个节点的上一个更新为新节点的指针
	l.length++            //链表长度+1
	return l.Len() - 1
}

func (l *LinkedList[T]) QueueJump(index int, value T) (err error) {
	if err = l.isOutOfRange(index); err != nil {
		return
	}
	tmpNode := l.findNode(index) //获取原本下标的节点
	node := &node[T]{
		prev:  tmpNode.prev, //为原本节点的上一个
		next:  tmpNode,      //为原本节点
		value: value,
	}
	node.prev.next = node //将上一个节点的下一个更新为新节点的指针
	node.next.prev = node //将下一个节点的上一个更新为新节点的指针
	l.length++            //链表长度+1
	return
}

func (l *LinkedList[T]) Delete(index int) (value T, err error) {
	if err = l.isOutOfRange(index); err != nil {
		return
	}
	tmpNode := l.findNode(index)     //获取原本下标的节点
	tmpNode.prev.next = tmpNode.next //将上一个节点的下一个 更新为下个顺位节点的指针
	tmpNode.next.prev = tmpNode.prev //将下一个节点的上一个 更新为前个顺位节点的指针
	l.length--                       //链表长度-1
	value = tmpNode.value
	return
}

func (l *LinkedList[T]) Len() int {
	return l.length
}

func (l *LinkedList[T]) Cap() int {
	return l.Len()
}

func (l *LinkedList[T]) Range(f func(index int, value T) error) error {
	tmp := l.head.next
	for i := 0; i < l.Len(); i++ {
		if err := f(i, tmp.value); err != nil {
			return err
		}
		tmp = tmp.next
	}
	return nil
}

func (l *LinkedList[T]) GetSlice() []T {
	s := make([]T, l.length)
	tmp := l.head.next
	for i := 0; i < l.length; i++ {
		s[i] = tmp.value
		tmp = tmp.next
	}
	return s
}
