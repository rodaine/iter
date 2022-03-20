package iter

type Ordered interface {
	Integer | float32 | float64
}

func min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T Ordered](iter Iterator[T]) (T, bool) {
	return iter.Reduce(min[T])
}

func Max[T Ordered](iter Iterator[T]) (T, bool) {
	return iter.Reduce(max[T])
}
