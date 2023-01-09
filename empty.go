package iter

// Empty returns an Iterator that is already finished, never returning an
// element.
func Empty[E any]() Iterator[E] { return FromCoreFunc(empty[E]) }

func empty[E any]() (next E, ok bool) { return }
