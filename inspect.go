package iter

type inspectCore[T any] struct {
	Core[T]
	Fn func(T)
}

func (ic inspectCore[T]) Next() (T, bool) {
	next, ok := ic.Core.Next()
	if ok {
		ic.Fn(next)
	}
	return next, ok
}

func (iter Iterator[T]) Inspect(fn func(T)) Iterator[T] {
	return FromCore[T](inspectCore[T]{
		Core: iter.core.core(),
		Fn:   fn,
	})
}
