package iter

// A Predicate function receives a value and returns true if the value matches
// the predicate and should be emitted.
type Predicate[T any] func(value T) (matches bool)

// A TryPredicate is a predicate that is fallible, returning an error if
// execution should short-circuit.
type TryPredicate[T any] func(value T) (matches bool, err error)

// Filter returns a new iterator that filters elements matching the provided
// Predicate.
func (iter Iterator[E]) Filter(pred Predicate[E]) Iterator[E] {
	return FromCore[E](filterCore[E]{
		Core: iter.core,
		pred: pred,
	})
}

type filterCore[E any] struct {
	Core[E]
	pred Predicate[E]
}

func (fc filterCore[E]) Next() (E, bool) {
	for next, ok := fc.Core.Next(); ok; next, ok = fc.Core.Next() {
		if fc.pred(next) {
			return next, true
		}
	}

	return empty[E]()
}
