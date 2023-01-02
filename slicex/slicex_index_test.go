package slicex

import (
	"testing"
)

func TestExcludeByIndex(t *testing.T) {
	arr := []int{1, 23, 42352, 523}
	index, el, err := ExcludeByIndex(arr, 3)
	t.Log(index, el, err)
}
