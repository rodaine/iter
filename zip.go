package iter

type Pair[T, U any] struct {
	A T
	B U
}

type zipCore[T, U any] struct {
	a Core[T]
	b Core[U]
}

func (z zipCore[T, U]) Next() (Pair[T, U], bool) {
	a, ok := z.a.Next()
	if !ok {
		return Pair[T, U]{}, false
	}

	b, ok := z.b.Next()
	if !ok {
		return Pair[T, U]{}, false
	}

	return Pair[T, U]{
		A: a,
		B: b,
	}, true
}

func Zip[A, B any](a Iterator[A], b Iterator[B]) Iterator[Pair[A, B]] {
	return FromCore[Pair[A, B]](zipCore[A, B]{
		a: a.core.core(),
		b: b.core.core(),
	})
}

func Enumerate[T any](iter Iterator[T]) Iterator[Pair[int, T]] {
	return Zip(CountUp[int](0), iter)
}
