package iter

// FromSlice returns an Iterator over the provided slice.
func FromSlice[E any](s []E) Iterator[E] {
	return FromCore[E](&sliceCore[E]{s: s})
}

// FromItems creates an Iterator from the provided items.
func FromItems[E any](items ...E) Iterator[E] {
	return FromSlice(items)
}

// ToSlice collects the elements of the Iterator into a slice.
func (iter Iterator[E]) ToSlice() []E {
	var out []E

	iter.ForEach(func(el E) {
		out = append(out, el)
	})

	return out
}

type sliceCore[E any] struct{ s []E }

func (si *sliceCore[E]) Next() (E, bool) {
	if len(si.s) == 0 {
		var zero E
		return zero, false
	}

	out := si.s[0]
	si.s = si.s[1:]
	return out, true
}
