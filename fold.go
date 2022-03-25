package iter

// A FoldFunc is an accumulator function takes the current state and the next
// value, and returns a new version of the state.
type FoldFunc[S, T any] func(state S, value T) S

// TryFoldFunc is a FoldFunc that may also return an error, short-circuiting
// accumulation.
type TryFoldFunc[S, T any] func(state S, value T) (S, error)

// Fold applies an accumulator on an Iterator's elements, reducing it to a
// final resulting value. Fold takes an Iterator, an initial value for the
// accumulator, and a FoldFunc that is applied to the current value of the
// accumulator and the next element of the Iterator. After all elements of the
// iterator are consumed, the final accumulator value is returned.
//
// Iterator.Reduce is a specialization of Fold, using the first element of the
// iterator as the initial state.
func Fold[S, E any](iter Iterator[E], initial S, fn FoldFunc[S, E]) S {
	state := initial
	iter.ForEach(func(el E) { state = fn(state, el) })
	return state
}

// TryFold is the same as Fold, but takes a TryFoldFunc which allows for
// fallible accumulation. The first error results in this function
// short-circuiting, returning the zero value for the accumulator state and the
// error.
func TryFold[S, E any](iter Iterator[E], initial S, fn TryFoldFunc[S, E]) (S, error) {
	state := initial
	var err error

	err = iter.TryForEach(func(el E) error {
		state, err = fn(state, el)
		return err
	})

	return state, err
}
