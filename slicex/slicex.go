package slicex

// ArrayHas 判断是否包含，支持可比较数据类型
func ArrayHas[T comparable](array []T, has T) bool {
	return ArrayHasFunc(array, has, func(src, has T) bool {
		return src == has
	})
}

// ArrayHasFunc 判断是否包含，支持自定义比较方式
func ArrayHasFunc[T any](array []T, has T, f func(src, has T) bool) bool {
	for _, arr := range array {
		if f(arr, has) {
			return true
		}
	}
	return false
}

// ArrayHasAny 判断是否完全包含，支持可比较数据类型
func ArrayHasAny[T comparable](array []T, has []T) bool {
	arr := make(map[T]struct{}, len(array))
	for _, v := range array {
		arr[v] = struct{}{}
	}

	count := 0
	for _, h := range has {
		if _, ok := arr[h]; ok {
			count++
		}
	}
	return count == len(has)
}

// ArrayHasAnyFunc 判断是否完全包含，支持自定义比较方式
func ArrayHasAnyFunc[T any](array []T, has []T, f func(src, has T) bool) bool {
	count := 0
	for _, h := range has {
		for _, a := range array {
			if f(a, h) {
				count++
			}
		}
	}
	return count == len(has)
}
