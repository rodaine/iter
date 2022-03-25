package iter

// Skip returns a new Iterator that advances the current iterator by n before
// returning the Next element.
func (iter Iterator[E]) Skip(n uint) Iterator[E] {
	return FromCore[E](&skipCore[E]{
		Iterator: iter,
		n:        n,
	})
}

// SkipWhile returns a new Iterator that will skip values matching the provided
// Predicate until the first element that does not match. After the first
// non-matching element, the predicate is no longer applied.
func (iter Iterator[E]) SkipWhile(pred Predicate[E]) Iterator[E] {
	return FromCore[E](&skipWhileCore[E]{
		Core: iter.core,
		pred: pred,
	})
}

type skipCore[E any] struct {
	Iterator[E]
	n uint
}

func (sc *skipCore[E]) Next() (E, bool) {
	if sc.n > 0 {
		sc.Iterator.AdvanceBy(sc.n)
		sc.n = 0
	}
	return sc.Iterator.Next()
}

type skipWhileCore[E any] struct {
	Core[E]
	pred Predicate[E]
	done bool
}

func (swc *skipWhileCore[E]) Next() (E, bool) {
	for next, ok := swc.Core.Next(); ok; next, ok = swc.Core.Next() {
		if swc.done = swc.done || !swc.pred(next); swc.done {
			return next, ok
		}
	}

	return empty[E]()
}
