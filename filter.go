package iter

type filterCore[T any] struct {
	Core[T]
	Fn func(T) bool
}

func (fc filterCore[T]) Next() (T, bool) {
	for next, ok := fc.Core.Next(); ok; next, ok = fc.Core.Next() {
		if fc.Fn(next) {
			return next, true
		}
	}

	return empty[T]()
}

func (iter Iterator[T]) Filter(fn func(T) bool) Iterator[T] {
	return FromCore[T](filterCore[T]{
		Core: iter.core.core(),
		Fn:   fn,
	})
}
