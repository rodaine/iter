package iter

// Chain returns a new Iterator which will first iterate over values from the
// original iter and then over values from the subsequent iterators in order.
func (iter Iterator[E]) Chain(iters ...Iterator[E]) Iterator[E] {
	if len(iters) == 0 {
		return iter
	}

	chain := make([]Core[E], 1+len(iters))
	chain[0] = iter.core
	for i, it := range iters {
		chain[1+i] = it.core
	}

	return FromCore[E](&chainCore[E]{chain: chain})
}

type chainCore[E any] struct {
	chain []Core[E]
}

func (c *chainCore[E]) Next() (E, bool) {
	for len(c.chain) != 0 {
		if next, ok := c.chain[0].Next(); ok {
			return next, ok
		}

		c.chain = c.chain[1:]
	}
	return empty[E]()
}
