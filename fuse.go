package iter

// Fuse creates an Iterator.Next that permanently ends after the first
// (_, false). After an Iterator.Next false,future calls may or may not yield a
// value and true. Fuse ensures that after a false is given, it will return
// false thereafter.
func (iter Iterator[E]) Fuse() Iterator[E] {
	return FromCore[E](&fuseCore[E]{
		Core: iter.core,
	})
}

type fuseCore[E any] struct {
	Core[E]
	done bool
}

func (fc *fuseCore[E]) Next() (E, bool) {
	if fc.done {
		return empty[E]()
	}

	if next, ok := fc.Core.Next(); ok {
		return next, ok
	}

	fc.done = true
	return empty[E]()
}
