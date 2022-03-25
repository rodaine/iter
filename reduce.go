package iter

// ReduceFunc is a FoldFunc that uses the same type for the state and element
// values.
type ReduceFunc[T any] FoldFunc[T, T]

// TryReduceFunc is a TryFoldFunc that uses the same type for the state and
// element values.
type TryReduceFunc[T any] TryFoldFunc[T, T]

// Reduce is a specialization of Fold that uses the first element of the
// Iterator as the initial accumulator value. The ok value is false if the
// iterator is empty.
func (iter Iterator[E]) Reduce(fn ReduceFunc[E]) (result E, ok bool) {
	initial, ok := iter.Next()
	if !ok {
		return empty[E]()
	}

	return Fold(iter, initial, FoldFunc[E, E](fn)), true
}

// TryReduce is a specialization of TryFold that uses the first element of the
// Iterator as the initial accumulator value. The ok value is false if the
// iterator is empty.
func (iter Iterator[E]) TryReduce(fn TryReduceFunc[E]) (result E, ok bool, err error) {
	initial, ok := iter.Next()
	if !ok {
		var zero E
		return zero, false, nil
	}

	out, err := TryFold(iter, initial, TryFoldFunc[E, E](fn))
	return out, true, err
}
