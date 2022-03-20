package iter

type filterMapCore[T, U any] struct {
	Core[T]
	Fn func(T) (U, bool)
}

func (fmc filterMapCore[T, U]) Next() (U, bool) {
	for next, ok := fmc.Core.Next(); ok; next, ok = fmc.Core.Next() {
		if out, ok := fmc.Fn(next); ok {
			return out, true
		}
	}

	return empty[U]()
}

func FilterMap[T, U any](iter Iterator[T], fn func(T) (U, bool)) Iterator[U] {
	return FromCore[U](filterMapCore[T, U]{
		Core: iter.core.core(),
		Fn:   fn,
	})
}
