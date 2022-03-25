package iter

// FilterMapFunc converts a value of type A into B, returning ok as true if the
// value should be emitted.
type FilterMapFunc[A, B any] func(from A) (to B, ok bool)

// FilterMap returns an iterator that both filters and maps. The returned
// iterator yields only yields value for which the supplied FilterMapFunc
// returns (value, true).
func FilterMap[A, B any](iter Iterator[A], fn FilterMapFunc[A, B]) Iterator[B] {
	return FromCore[B](filterMapCore[A, B]{
		Core: iter.core,
		fn:   fn,
	})
}

type filterMapCore[A, B any] struct {
	Core[A]
	fn FilterMapFunc[A, B]
}

func (fmc filterMapCore[_, B]) Next() (B, bool) {
	for next, ok := fmc.Core.Next(); ok; next, ok = fmc.Core.Next() {
		if out, ok := fmc.fn(next); ok {
			return out, true
		}
	}

	return empty[B]()
}
