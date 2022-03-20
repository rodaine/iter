package iter

type takeCore[T any] struct {
	Core[T]
	n uint
}

func (tc *takeCore[T]) Next() (T, bool) {
	if tc.n == 0 {
		return empty[T]()
	}

	tc.n--
	return tc.Core.Next()
}

func (iter Iterator[T]) Take(n uint) Iterator[T] {
	return FromCore[T](&takeCore[T]{
		Core: iter.core.core(),
		n:    n,
	})
}

type takeWhileCore[T any] struct {
	Core[T]
	Fn   func(T) bool
	done bool
}

func (twc *takeWhileCore[T]) Next() (T, bool) {
	if twc.done {
		return empty[T]()
	}

	next, ok := twc.Core.Next()
	if twc.done = !twc.Fn(next); twc.done {
		return empty[T]()
	}

	return next, ok
}

func (iter Iterator[T]) TakeWhile(fn func(T) bool) Iterator[T] {
	return FromCore[T](&takeWhileCore[T]{
		Core: iter.core.core(),
		Fn:   fn,
	})
}
