package iter

// Pair is a "tuple" of two values of any types.
type Pair[A, B any] struct {
	A A
	B B
}

// Zip merges two Iterator into a single iterator producing Pair of the values
// emitted by the original iterators in lock-step. The returned iterator is
// finished if either of the parent iterators completes.
func Zip[E1, E2 any](a Iterator[E1], b Iterator[E2]) Iterator[Pair[E1, E2]] {
	return FromCore[Pair[E1, E2]](zipCore[E1, E2]{
		a: a.core,
		b: b.core,
	})
}

// Enumerate zips the elements of the provided Iterator with a zero-indexed
// counter of each element.
func Enumerate[E any](iter Iterator[E]) Iterator[Pair[int, E]] {
	return Zip(CountUpBy[int](0, 1), iter)
}

type zipCore[E1, E2 any] struct {
	a Core[E1]
	b Core[E2]
}

func (z zipCore[E1, E2]) Next() (Pair[E1, E2], bool) {
	a, ok := z.a.Next()
	if !ok {
		return Pair[E1, E2]{}, false
	}

	b, ok := z.b.Next()
	if !ok {
		return Pair[E1, E2]{}, false
	}

	return Pair[E1, E2]{
		A: a,
		B: b,
	}, true
}
