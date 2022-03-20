package iter

type mapCore[T, U any] struct {
	Core[T]
	Fn func(t T) U
}

func (mc mapCore[T, U]) Next() (U, bool) {
	next, ok := mc.Core.Next()
	if !ok {
		return empty[U]()
	}

	return mc.Fn(next), true
}

func Map[T, U any](iter Iterator[T], fn func(T) U) Iterator[U] {
	return FromCore[U](mapCore[T, U]{
		Core: iter.core.core(),
		Fn:   fn,
	})
}

type mapWhileCore[T, U any] struct {
	Core[T]
	Fn   func(t T) (U, bool)
	done bool
}

func (mwc *mapWhileCore[T, U]) Next() (U, bool) {
	if mwc.done {
		return empty[U]()
	}

	next, ok := mwc.Core.Next()
	if !ok {
		mwc.done = true
		return empty[U]()
	}

	out, ok := mwc.Fn(next)
	if mwc.done = !ok; mwc.done {
		return empty[U]()
	}

	return out, true
}

func MapWhile[T, U any](iter Iterator[T], fn func(T) (U, bool)) Iterator[U] {
	return FromCore[U](&mapWhileCore[T, U]{
		Core: iter.core.core(),
		Fn:   fn,
	})
}
