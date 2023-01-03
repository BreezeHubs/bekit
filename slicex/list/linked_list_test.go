package list

import (
	"fmt"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	list := NewLinkedList[int]()
	for i := 0; i < 5; i++ {
		list.Append(i)
	}

	list.Range(func(index int, value int) error {
		fmt.Println(index, value)
		return nil
	})
}
