package iter

// A ScanFunc receives a pointer to the current state and the next value,
// returning a result value and ok indicating if the iterator should
// continue.
type ScanFunc[S, A, B any] func(state *S, from A) (to B, ok bool)

// Scan behaves similar to Fold and MapWhile, but maintains internal state and
// produces a new Iterator. A pointer to the initial state is passed to the
// provided ScanFunc for each element on the Iterator. This state can be
// mutated within the ScanFunc to share state between iterations.
func Scan[S, A, B any](iter Iterator[A], initial S, fn ScanFunc[S, A, B]) Iterator[B] {
	return FromCore[B](scanCore[S, A, B]{
		Core:  iter.core,
		state: &initial,
		fn:    fn,
	})
}

type scanCore[S, A, B any] struct {
	Core[A]
	state *S
	fn    ScanFunc[S, A, B]
}

func (sc scanCore[S, A, B]) Next() (B, bool) {
	next, ok := sc.Core.Next()
	if !ok {
		return empty[B]()
	}

	return sc.fn(sc.state, next)
}
