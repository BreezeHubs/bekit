package list

import (
	"fmt"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	list := NewLinkedList[int]()
	for i := 1; i <= 5; i++ {
		list.Append(i)
	}
	list.Range(func(index int, value int) error {
		fmt.Println(index, " => ", value)
		return nil
	})

	fmt.Println("插队2：", list.QueueJump(2, 999), list.GetSlice())
	fmt.Println("插队5：", list.QueueJump(5, 998), list.GetSlice())

	val, err := list.Delete(4)
	fmt.Println("删除4：", val, err)

	list.Range(func(index int, value int) error {
		fmt.Print(index, "=>", value, " | ")
		return nil
	})

	fmt.Println("\n", list.GetSlice())
}
