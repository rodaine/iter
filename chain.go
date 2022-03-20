package iter

type chainCore[T any] struct {
	chain []Core[T]
}

func (c *chainCore[T]) Next() (T, bool) {
	if len(c.chain) == 0 {
		return empty[T]()
	}

	if next, ok := c.chain[0].Next(); ok {
		return next, ok
	}

	c.chain = c.chain[1:]
	return c.Next()
}

func (iter Iterator[T]) Chain(iters ...Iterator[T]) Iterator[T] {
	chain := make([]Core[T], 1+len(iters))
	chain[0] = iter.core.core()
	for i, it := range iters {
		chain[1+i] = it.core.core()
	}

	return FromCore[T](&chainCore[T]{chain: chain})
}
