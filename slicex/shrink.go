package slicex

// 缩容计算
func calcCapacity(c, l int) (int, bool) {
	//cap 大于64才需要进行处理
	if c <= 64 {
		return c, false
	}

	//容量为大小的 2+倍 时，以0.625缩容
	if c > 2048 && (c/l >= 2) {
		factor := 0.625
		return int(float32(c) * float32(factor)), true
	}

	//容量为大小的 4+倍 时，以一半缩容
	if c <= 2048 && (c/l >= 4) {
		return c / 2, true
	}

	return c, false
}

// ArrayShrink 缩容成新数组
func ArrayShrink[T any](src []T) []T {
	n, changed := calcCapacity(cap(src), len(src))
	if !changed {
		return src
	}
	s := make([]T, 0, n)
	s = append(s, src...)
	return s
}
