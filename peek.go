package iter

// A PeekIterator is an Iterator that also supports the Peek operation. A
// PeekIterator may be created directly via FromPeekCore or by upgrading an
// existing Iterator via Iterator.Peekable.
type PeekIterator[E any] struct {
	Iterator[E]
	core PeekCore[E]
}

// Peekable converts the current Iterator into a PeekIterator.
func (iter Iterator[E]) Peekable() PeekIterator[E] {
	return FromPeekCore[E](&peekCore[E]{
		Core: iter.core,
	})
}

// Peekable on a PeekIterator is as noop.
func (pi PeekIterator[E]) Peekable() PeekIterator[E] {
	return pi
}

// Peek returns the next value without advancing the Iterator.
func (pi PeekIterator[E]) Peek() (next E, ok bool) {
	return pi.core.Peek()
}

// A PeekCore extends a Core to also provide Peek functionality. Custom
// implementations of a Core may expose specialized (e.g., more efficient) Peek
// methods via this interface; otherwise, Iterator.Peekable will upgrade any
// iterator to a PeekIterator.
type PeekCore[E any] interface {
	Core[E]

	// Peek looks ahead to the next element, without advancing the Core.
	Peek() (E, bool)
}

// FromPeekCore converts a PeekCore into a PeekIterator.
func FromPeekCore[E any](core PeekCore[E]) PeekIterator[E] {
	return PeekIterator[E]{
		core:     core,
		Iterator: FromCore[E](core),
	}
}

type peekCore[E any] struct {
	Core[E]
	peeked bool
	peek   E
}

func (p *peekCore[E]) Next() (next E, ok bool) {
	if p.peeked {
		p.peeked = false
		return p.peek, true
	}

	return p.Core.Next()
}

func (p *peekCore[E]) Peek() (next E, ok bool) {
	if p.peeked {
		return p.peek, true
	}

	p.peek, p.peeked = p.Core.Next()
	return p.peek, p.peeked
}
