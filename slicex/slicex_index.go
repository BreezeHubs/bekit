package slicex

import "errors"

// IndexError 下标错误
var IndexError = errors.New("下标错误")

// FindByIndex 返回数据对应下标的数据，会判断是否存在下标
func FindByIndex[T any](array []T, idx int) (val T, err error) {
	if idx < 0 || idx >= len(array) {
		err = IndexError
		return
	}
	val = array[idx]
	return
}

// ExcludeByIndex 返回去除对应下标的数组，返回新数组、对应下标数据，会判断是否存在下标
func ExcludeByIndex[T any](array []T, idx int) (res []T, val T, err error) {
	if idx < 0 || idx >= len(array) {
		err = IndexError
		return
	}

	res = append(array[:idx], array[idx+1:]...)
	val = array[idx]
	return
}
