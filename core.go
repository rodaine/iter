package iter

type Core[T any] interface {
	Next() (T, bool)
}

func FromCore[T any](core Core[T]) Iterator[T] {
	return Iterator[T]{&peekCore[T]{
		Core: core,
	}}
}

type CoreFunc[T any] func() (T, bool)

func (cf CoreFunc[T]) Next() (T, bool) { return cf() }

func FromCoreFunc[T any](fn CoreFunc[T]) Iterator[T] {
	return FromCore[T](fn)
}
