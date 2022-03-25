package iter

// An Iterator allows for performing streaming operations against elements of a
// collection or other source of data. At its most fundamental, an Iterator is
// defined by Core.Next, with all other functionality layered upon it. Custom
// implementations need only implement Core, using FromCore to expose the rest
// of the Iterator functionality.
type Iterator[E any] struct {
	core Core[E]
}

// Next advances the iterator and returns the next element. The ok value is set
// to false when the iteration is finished. Individual iterator implementations
// may choose to resume iteration, and so calling Next again may or may not
// eventually start returning values again at some point.
func (iter Iterator[E]) Next() (next E, ok bool) {
	return iter.core.Next()
}

// Count calls Iterator.Next repeatedly until the first time ok is false, returning
// the number of elements it saw.
func (iter Iterator[_]) Count() (ct uint) {
	for _, ok := iter.Next(); ok; _, ok = iter.Next() {
		ct++
	}
	return ct
}

// Last calls Iterator.Next repeatedly until the first time ok is false,
// returning the last element seen. The ok value is false if the Iterator is
// already finished.
func (iter Iterator[E]) Last() (last E, ok bool) {
	for out, cont := iter.Next(); cont; out, cont = iter.Next() {
		last = out
		ok = true
	}

	return last, ok
}

// AdvanceBy will skip up to n elements by calling Iterator.Next up to n times
// or until ok is false. If the iterator has less than n items, the returned ct
// will be the number of elements skipped with the ok being false. Otherwise,
// ct will equal n and ok will be true.
func (iter Iterator[_]) AdvanceBy(n uint) (ct uint, ok bool) {
	for i := uint(0); i < n; i++ {
		if _, ok = iter.Next(); !ok {
			return i, false
		}
	}
	return n, true
}

// Nth returns the zero-indexed element of the iterator. This method will
// return ok as false if n is greater than or equal to the length of the
// iterator.
func (iter Iterator[E]) Nth(n uint) (nth E, ok bool) {
	for i := uint(0); i <= n; i++ {
		if nth, ok = iter.Next(); !ok {
			return nth, ok
		}
	}
	return nth, ok
}

// ForEach consumes the Iterator, calling the provided function on each element
// of the Iterator.
func (iter Iterator[E]) ForEach(fn func(E)) {
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		fn(next)
	}
}

// TryForEach performs the same behavior as ForEach, but if the provided
// function returns an error, execution stops early and the error is returned.
func (iter Iterator[E]) TryForEach(fn func(E) error) error {
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		if err := fn(next); err != nil {
			return err
		}
	}
	return nil
}

// Partition applies the provided Predicate to the iterator, returning two
// slices of elements matching or not matching, respectively.
func (iter Iterator[E]) Partition(pred Predicate[E]) (trues []E, falses []E) {
	for el, ok := iter.Next(); ok; el, ok = iter.Next() {
		if pred(el) {
			trues = append(trues, el)
		} else {
			falses = append(falses, el)
		}
	}

	return
}

// All returns true if all elements of the Iterator match the provided
// Predicate.
func (iter Iterator[E]) All(pred Predicate[E]) bool {
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		if !pred(next) {
			return false
		}
	}

	return true
}

// Any returns true if at least one element of the Iterator matches the
// provided Predicate.
func (iter Iterator[E]) Any(pred Predicate[E]) bool {
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		if pred(next) {
			return true
		}
	}
	return false
}

// Find returns the first element of the Iterator that matches the provided
// Predicate.
func (iter Iterator[E]) Find(pred Predicate[E]) (match E, ok bool) {
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		if pred(next) {
			return next, true
		}
	}

	return empty[E]()
}

// TryFind behaves the same as Find, but will stop execution if the provided
// TryPredicate returns an error.
func (iter Iterator[E]) TryFind(pred TryPredicate[E]) (match E, ok bool, err error) {
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		if match, err := pred(next); err != nil {
			var zero E
			return zero, false, err
		} else if match {
			return next, true, nil
		}
	}

	var zero E
	return zero, false, nil
}

// Position returns the zero-based index of the element in the Iterator that
// matches the provided Predicate.
func (iter Iterator[E]) Position(pred Predicate[E]) (idx uint, ok bool) {
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		if pred(next) {
			return idx, true
		}
		idx++
	}
	return empty[uint]()
}
