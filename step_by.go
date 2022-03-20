package iter

type stepByCore[T any] struct {
	iter    Iterator[T]
	nth     uint
	started bool
}

func (sbc *stepByCore[T]) Next() (T, bool) {
	if !sbc.started {
		sbc.started = true
		return sbc.iter.Next()
	}

	return sbc.iter.Nth(sbc.nth)
}

func (iter Iterator[T]) StepBy(n uint) Iterator[T] {
	return FromCore[T](&stepByCore[T]{
		iter: iter,
		nth:  n - 1,
	})
}
