package iter

import "bufio"

// FromStringScanner converts a bufio.Scanner into an Iterator, returning the
// strings emitted by bufio.Scanner.Text. Note that the iterator does not
// produce errors, which would need to be independently checked on the
// bufio.Scanner after execution.
func FromStringScanner(sc *bufio.Scanner) Iterator[string] {
	return FromCore[string](scannerCore[string]{
		sc:   sc,
		next: sc.Text,
	})
}

// FromBytesScanner converts a bufio.Scanner into an Iterator, returning a copy
// of the []byte emitted by bufio.Scanner.Bytes. Note that the iterator does
// not produce errors, which would need to be independently checked on the
// bufio.Scanner after execution.
func FromBytesScanner(sc *bufio.Scanner) Iterator[[]byte] {
	return FromCore[[]byte](scannerCore[[]byte]{
		sc: sc,
		next: func() []byte {
			b := sc.Bytes()
			out := make([]byte, len(b))
			copy(out, b)
			return out
		},
	})
}

type scannerCore[E string | []byte] struct {
	sc   *bufio.Scanner
	next func() E
}

func (s scannerCore[E]) Next() (E, bool) {
	if !s.sc.Scan() {
		return empty[E]()
	}
	return s.next(), true
}
