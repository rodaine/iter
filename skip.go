package iter

type skipCore[T any] struct {
	Iterator[T]
	n uint
}

func (sc *skipCore[T]) Next() (T, bool) {
	if sc.n > 0 {
		sc.Iterator.AdvanceBy(sc.n)
		sc.n = 0
	}
	return sc.Iterator.Next()
}

func (iter Iterator[T]) Skip(n uint) Iterator[T] {
	return FromCore[T](&skipCore[T]{
		Iterator: iter,
		n:        n,
	})
}

type skipWhileCore[T any] struct {
	Core[T]
	Fn   func(T) bool
	done bool
}

func (swc *skipWhileCore[T]) Next() (T, bool) {
	for next, ok := swc.Core.Next(); ok; next, ok = swc.Core.Next() {
		if swc.done = swc.done || !swc.Fn(next); swc.done {
			return next, ok
		}
	}

	return empty[T]()
}

func (iter Iterator[T]) SkipWhile(fn func(T) bool) Iterator[T] {
	return FromCore[T](&skipWhileCore[T]{
		Core: iter.core.core(),
		Fn:   fn,
	})
}
