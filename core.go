package iter

// Core represent the basic functionality of an Iterator that enables all other
// features exposed on the type. Custom implementations typically only need to
// implement a Core (or PeekCore).
type Core[E any] interface {
	// Next advances the iterator and returns the next element. The ok value is
	// set to false when the iteration is finished. Individual iterator
	// implementations may choose to resume iteration, and so calling Next
	// again may or may not eventually start returning values again at some
	// point.
	Next() (next E, ok bool)
}

// FromCore upgrades a Core type into a full Iterator.
func FromCore[E any](core Core[E]) Iterator[E] {
	return Iterator[E]{core: core}
}

// CoreFunc is a convenience type for implementing Core for a function that
// matches the signature of Core.Next.
type CoreFunc[E any] func() (next E, ok bool)

// Next satisfies the Core interface.
func (cf CoreFunc[E]) Next() (next E, ok bool) { return cf() }

// FromCoreFunc creates an Iterator directly from a CoreFunc.
func FromCoreFunc[E any](fn CoreFunc[E]) Iterator[E] {
	return FromCore[E](fn)
}
