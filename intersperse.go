package iter

// Intersperse creates a new Iterator which places the provided separator
// between subsequent elements.
func (iter Iterator[E]) Intersperse(sep E) Iterator[E] {
	return iter.IntersperseWith(func() E { return sep })
}

// IntersperseWith behaves the same as Iterator.Intersperse, but uses the
// provided function to generate the separator.
func (iter Iterator[E]) IntersperseWith(sepFn func() E) Iterator[E] {
	return FromCore[E](&intersperseCore[E]{
		core: iter.Peekable(),
		sep:  sepFn,
	})
}

type intersperseCore[E any] struct {
	core    PeekCore[E]
	sep     func() E
	sepNext bool
}

func (i *intersperseCore[E]) Next() (E, bool) {
	if i.sepNext {
		i.sepNext = false
		return i.sep(), true
	}

	next, ok := i.core.Next()
	_, i.sepNext = i.core.Peek()
	return next, ok
}
