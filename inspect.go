package iter

// Inspect does something with each element of the Iterator before passing it
// on. Inspect is typically used as a debugging tool to check out the internals
// of a chain of iterator steps at various points in the pipeline.
func (iter Iterator[E]) Inspect(fn func(E)) Iterator[E] {
	return FromCore[E](inspectCore[E]{
		Core: iter.core,
		fn:   fn,
	})
}

type inspectCore[E any] struct {
	Core[E]
	fn func(E)
}

func (ic inspectCore[E]) Next() (E, bool) {
	next, ok := ic.Core.Next()
	if ok {
		ic.fn(next)
	}
	return next, ok
}
