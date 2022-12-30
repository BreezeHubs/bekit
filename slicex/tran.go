package slicex

func ArrayToMap[T any](array []T) map[int]T {
	m := make(map[int]T, len(array))
	for i, v := range array {
		m[i] = v
	}
	return m
}

func MapToArray[T any, K comparable](m map[K]T) []T {
	array := []T{}
	for _, v := range m {
		array = append(array, v)
	}
	return array
}
