package slicex

import "errors"

type arrayT interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | string
}

func For[T arrayT](arr []T, f func(T, int) any) ([]T, error) {
	var newArr []T
	for i, a := range arr {
		v := f(a, i)
		switch v.(type) {
		case T:
			newArr = append(newArr, v.(T))
		case bool:
		default:
			return nil, errors.New("func has a return type that is not allowed")
		}
	}
	return newArr, nil
}
