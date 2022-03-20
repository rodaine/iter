package iter

type scanCore[S, T, U any] struct {
	Core[T]
	State *S
	Fn    func(*S, T) (U, bool)
}

func (sc scanCore[S, T, U]) Next() (U, bool) {
	next, ok := sc.Core.Next()
	if !ok {
		return empty[U]()
	}

	return sc.Fn(sc.State, next)
}

func Scan[S, T, U any](iter Iterator[T], initialState S, fn func(*S, T) (U, bool)) Iterator[U] {
	return FromCore[U](scanCore[S, T, U]{
		Core:  iter.core.core(),
		State: &initialState,
		Fn:    fn,
	})
}
