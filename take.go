package iter

// Take returns a new Iterator that limits the number of elements returned to
// the specified number.
func (iter Iterator[E]) Take(n uint) Iterator[E] {
	return FromCore[E](&takeCore[E]{
		Core: iter.core,
		n:    n,
	})
}

// TakeWhile returns a new Iterator that emits values that match the provided
// Predicate, stopping after the first non-match.
func (iter Iterator[E]) TakeWhile(pred Predicate[E]) Iterator[E] {
	return FromCore[E](&takeWhileCore[E]{
		Core: iter.core,
		pred: pred,
	})
}

type takeCore[E any] struct {
	Core[E]
	n uint
}

func (tc *takeCore[E]) Next() (E, bool) {
	if tc.n == 0 {
		return empty[E]()
	}

	tc.n--
	return tc.Core.Next()
}

type takeWhileCore[E any] struct {
	Core[E]
	pred Predicate[E]
	done bool
}

func (twc *takeWhileCore[E]) Next() (E, bool) {
	if twc.done {
		return empty[E]()
	}

	next, ok := twc.Core.Next()
	if twc.done = !twc.pred(next); twc.done {
		return empty[E]()
	}

	return next, ok
}
