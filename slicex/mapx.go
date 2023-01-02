package slicex

// MapHas 判断是否包含，支持可比较数据类型
func MapHas[T comparable, K comparable](array map[K]T, has T) bool {
	return MapHasFunc(array, has, func(src, has T) bool {
		return src == has
	})
}

// MapHasFunc 判断是否包含，支持自定义比较方式
func MapHasFunc[T any, K comparable](array map[K]T, has T, f func(src, has T) bool) bool {
	for _, arr := range array {
		if f(arr, has) {
			return true
		}
	}
	return false
}

// MapHasAny 判断是否完全包含，支持可比较数据类型
func MapHasAny[T comparable, K comparable](m map[K]T, has []T) bool {
	arr := make(map[T]struct{}, len(m))
	for _, v := range m {
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

// MapHasAnyFunc 判断是否完全包含，支持自定义比较方式
func MapHasAnyFunc[T any, K comparable](m map[K]T, has []T, f func(src, has T) bool) bool {
	count := 0
	for _, h := range has {
		for _, a := range m {
			if f(a, h) {
				count++
			}
		}
	}
	return count == len(has)
}
