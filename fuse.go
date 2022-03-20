package iter

type fuseCore[T any] struct {
	Core[T]
	done bool
}

func (fc *fuseCore[T]) Next() (T, bool) {
	if fc.done {
		return empty[T]()
	}

	if next, ok := fc.Core.Next(); ok {
		return next, ok
	}

	fc.done = true
	return empty[T]()
}

func (iter Iterator[T]) Fuse() Iterator[T] {
	return FromCore[T](&fuseCore[T]{
		Core: iter.core.core(),
	})
}
