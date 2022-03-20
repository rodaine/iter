package iter

type peekCore[T any] struct {
	Core[T]
	peeked bool
	peek   T
}

func (p *peekCore[T]) Next() (T, bool) {
	if p.peeked {
		p.peeked = false
		return p.peek, true
	}

	return p.Core.Next()
}

func (p *peekCore[T]) Peek() (T, bool) {
	if p.peeked {
		return p.peek, true
	}

	p.peek, p.peeked = p.Core.Next()
	return p.peek, p.peeked
}

func (p *peekCore[T]) core() Core[T] {
	return p
}
