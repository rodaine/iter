package iter

// A FlatMapFunc is a MapFunc that emits an Iterator instead as its output value.
type FlatMapFunc[A, B any] MapFunc[A, Iterator[B]]

// FlatMap creates an iterator that works similar to Map, but flattens Iterators.
//
// Map is very useful, but only when the MapFunc produces values. If it
// produces an iterator instead, there's an extra layer of indirection. FlatMap
// removes that extra layer internally.
func FlatMap[A, B any](iter Iterator[A], fn FlatMapFunc[A, B]) Iterator[B] {
	return FromCore[B](&flatMapCore[A, B]{
		Core: iter.core,
		fn:   fn,
	})
}

type flatMapCore[A, B any] struct {
	Core[A]
	fn  FlatMapFunc[A, B]
	cur Core[B]
}

func (fmc *flatMapCore[_, B]) Next() (B, bool) {
	for {
		if fmc.cur == nil {
			next, ok := fmc.Core.Next()
			if !ok {
				return empty[B]()
			}
			fmc.cur = fmc.fn(next).core
		}

		if next, ok := fmc.cur.Next(); ok {
			return next, ok
		}

		fmc.cur = nil
	}
}
