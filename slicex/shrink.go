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
// 在必要的时候执行缩容，其算法为：
// - 如果容量 > 2048，并且长度小于容量一半，那么就会缩容为原本的 5/8
// - 如果容量 (64, 2048]，如果长度是容量的 1/4，那么就会缩容为原本的一半
// - 如果此时容量 <= 64，那么我们将不会执行缩容。在容量很小的情况下，浪费的内存很少，所以没必要消耗 CPU去执行缩容
func ArrayShrink[T any](src []T) []T {
	n, changed := calcCapacity(cap(src), len(src))
	if !changed {
		return src
	}
	s := make([]T, 0, n)
	s = append(s, src...)
	return s
}
