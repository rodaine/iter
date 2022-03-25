package iter

// MapFunc converts a value of type A into a value of type B.
type MapFunc[A, B any] func(from A) (to B)

// Map converts an Iterator of one type, converting it into an Iterator of
// another, using the provided MapFunc.
func Map[A, B any](iter Iterator[A], fn MapFunc[A, B]) Iterator[B] {
	return FromCore[B](mapCore[A, B]{
		Core: iter.core,
		fn:   fn,
	})
}

// MapWhile behaves like Map, but will stop emitting values if the provided
// FilterMapFunc returns false.
func MapWhile[A, B any](iter Iterator[A], fn FilterMapFunc[A, B]) Iterator[B] {
	return FromCore[B](&mapWhileCore[A, B]{
		Core: iter.core,
		Fn:   fn,
	})
}

type mapCore[A, B any] struct {
	Core[A]
	fn MapFunc[A, B]
}

func (mc mapCore[A, B]) Next() (B, bool) {
	next, ok := mc.Core.Next()
	if !ok {
		return empty[B]()
	}

	return mc.fn(next), true
}

type mapWhileCore[A, B any] struct {
	Core[A]
	Fn   FilterMapFunc[A, B]
	done bool
}

func (mwc *mapWhileCore[A, B]) Next() (B, bool) {
	if mwc.done {
		return empty[B]()
	}

	next, ok := mwc.Core.Next()
	if !ok {
		mwc.done = true
		return empty[B]()
	}

	out, ok := mwc.Fn(next)
	if mwc.done = !ok; mwc.done {
		return empty[B]()
	}

	return out, true
}
