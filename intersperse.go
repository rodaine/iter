package iter

type intersperseCore[T any] struct {
	core    *peekCore[T]
	sep     func() T
	sepNext bool
}

func (i *intersperseCore[T]) Next() (T, bool) {
	if i.sepNext {
		i.sepNext = false
		return i.sep(), true
	}

	next, ok := i.core.Next()
	_, i.sepNext = i.core.Peek()
	return next, ok
}

func (iter Iterator[T]) Intersperse(sep T) Iterator[T] {
	return iter.IntersperseWith(func() T { return sep })
}

func (iter Iterator[T]) IntersperseWith(fn func() T) Iterator[T] {
	return FromCore[T](&intersperseCore[T]{
		core: iter.core,
		sep:  fn,
	})
}
