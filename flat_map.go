package iter

type flatMapCore[T, U any] struct {
	Core[T]
	Fn  func(T) Iterator[U]
	Cur Core[U]
}

func (fmc *flatMapCore[T, U]) Next() (U, bool) {
	if fmc.Cur == nil {
		next, ok := fmc.Core.Next()
		if !ok {
			return empty[U]()
		}
		fmc.Cur = fmc.Fn(next).core.core()
	}

	if next, ok := fmc.Cur.Next(); ok {
		return next, ok
	}

	fmc.Cur = nil
	return fmc.Next()
}

func FlatMap[T, U any](iter Iterator[T], fn func(T) Iterator[U]) Iterator[U] {
	return FromCore[U](&flatMapCore[T, U]{
		Core: iter.core.core(),
		Fn:   fn,
	})
}
