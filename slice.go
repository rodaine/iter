package iter

type sliceCore[T any] struct{ s []T }

func (si *sliceCore[T]) Next() (T, bool) {
	if len(si.s) == 0 {
		var zero T
		return zero, false
	}

	out := si.s[0]
	si.s = si.s[1:]
	return out, true
}

func FromSlice[T any](s []T) Iterator[T] {
	return FromCore[T](&sliceCore[T]{s: s})
}

func FromItems[T any](s ...T) Iterator[T] {
	return FromSlice(s)
}

func (iter Iterator[T]) ToSlice() []T {
	var out []T

	iter.ForEach(func(el T) {
		out = append(out, el)
	})

	return out
}
