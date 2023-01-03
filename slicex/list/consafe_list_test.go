package list

import (
	"fmt"
	"testing"
)

func TestConSafeArrayList(t *testing.T) {
	list := NewConSafeArrayList[int](5)

	for i := 0; i < 5; i++ {
		list.Append(i)
		fmt.Println("新增：", list)
	}

	err := list.QueueJump(2, 3)
	fmt.Println("插队2,3：", err, list)

	err = list.Set(5, 9)
	fmt.Println("修改0,9：", err, list)

	list.Append(999)
	fmt.Println("追加：", list)

	tv, err := list.Delete(3)
	fmt.Println("删除3：", tv, err, list)

	fmt.Println("长度：", list.Len())
	fmt.Println("容量：", list.Cap())

	list.Range(func(index int, value int) error {
		fmt.Println(index, value)
		return nil
	})

	fmt.Println(list.GetSlice())
}
