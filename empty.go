package iter

func Empty[T any]() Iterator[T] { return FromCoreFunc[T](empty[T]) }

func empty[T any]() (next T, ok bool) { return }
