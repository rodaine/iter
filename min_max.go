package iter

import "golang.org/x/exp/constraints"

// KeyFunc converts a value of any type into a constraints.Ordered value that
// can be used to compare values.
type KeyFunc[T any, K constraints.Ordered] MapFunc[T, K]

// LessFunc compares two values, returning true if the first value is less than
// the second.
type LessFunc[T any] func(a, b T) (aLessThanB bool)

// Min consumes the provided Iterator, returning the minimum value emitted. The
// ok value will be false if the iterator is empty.
func Min[E constraints.Ordered](iter Iterator[E]) (min E, ok bool) {
	return iter.Reduce(minimum[E])
}

// MinBy consumes the Iterator, returning the minimum value determined by the
// specified LessFunc.
func (iter Iterator[E]) MinBy(fn LessFunc[E]) (min E, ok bool) {
	return iter.Reduce(func(a, b E) E {
		if fn(a, b) {
			return a
		}
		return b
	})
}

// MinByKey consumes the provided Iterator, returning the minimum value
// described by the value returned by the specified KeyFunc.
func MinByKey[E any, K constraints.Ordered](iter Iterator[E], fn KeyFunc[E, K]) (E, bool) {
	return iter.MinBy(func(a, b E) bool {
		return fn(a) < fn(b)
	})
}

// Max consumes the provided Iterator, returning the maximum value emitted. The
// ok value will be false if the iterator is empty.
func Max[E constraints.Ordered](iter Iterator[E]) (max E, ok bool) {
	return iter.Reduce(maximum[E])
}

// MaxBy consumes the provided Iterator, returning the maximum value determined
// by the specified LessFunc.
func (iter Iterator[E]) MaxBy(fn LessFunc[E]) (E, bool) {
	return iter.Reduce(func(a, b E) E {
		if !fn(a, b) {
			return a
		}
		return b
	})
}

// MaxByKey consumes the provided Iterator, returning the maximum value
// described by the value returned by the specified KeyFunc.
func MaxByKey[E any, K constraints.Ordered](iter Iterator[E], fn KeyFunc[E, K]) (E, bool) {
	return iter.MaxBy(func(a, b E) bool {
		return fn(a) < fn(b)
	})
}

func minimum[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func maximum[T constraints.Ordered](a, b T) T {
	if a >= b {
		return a
	}
	return b
}
