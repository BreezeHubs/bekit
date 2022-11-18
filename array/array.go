package array

import "bekit"

func Sum[T bekit.Number](ts []T) T {
	var res T
	for _, t := range ts {
		res += t
	}
	return res
}

func Max[T bekit.Number](ts []T) T {
	var res T
	for _, t := range ts {
		res += t
	}
	return res
}
