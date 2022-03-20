package iter

func (iter Iterator[T]) Reduce(fn func(T, T) T) (T, bool) {
	initial, ok := iter.Next()
	if !ok {
		return empty[T]()
	}

	return Fold(iter, initial, fn), true
}

func (iter Iterator[T]) TryReduce(fn func(T, T) (T, error)) (T, bool, error) {
	initial, ok := iter.Next()
	if !ok {
		var zero T
		return zero, false, nil
	}

	out, err := TryFold(iter, initial, fn)
	return out, true, err
}
