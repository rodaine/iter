package iter

func FromChan[T any](c <-chan T) Iterator[T] {
	return FromCoreFunc[T](func() (T, bool) {
		out, more := <-c
		return out, more
	})
}
