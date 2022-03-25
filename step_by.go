package iter

// StepBy returns a new iterator that returns every nth (one-indexed) element.
// The first element in the iterator is always returned, followed by every nth
// element thereafter. If n is equal to one, this method is a noop.
func (iter Iterator[E]) StepBy(n uint) Iterator[E] {
	if n <= 1 {
		return iter
	}

	return FromCore[E](&stepByCore[E]{
		iter: iter,
		nth:  n - 1,
	})
}

type stepByCore[E any] struct {
	iter    Iterator[E]
	nth     uint
	started bool
}

func (sbc *stepByCore[E]) Next() (E, bool) {
	if !sbc.started {
		sbc.started = true
		return sbc.iter.Next()
	}

	return sbc.iter.Nth(sbc.nth)
}
