package iter

func Fold[T, U any](iter Iterator[T], initial U, fn func(U, T) U) U {
	state := initial
	iter.ForEach(func(el T) { state = fn(state, el) })
	return state
}

func TryFold[T, U any](iter Iterator[T], initial U, fn func(U, T) (U, error)) (U, error) {
	state := initial
	var err error

	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		if state, err = fn(state, next); err != nil {
			var zero U
			return zero, err
		}
	}
	
	return state, nil
}
