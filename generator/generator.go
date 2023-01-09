package generator

import (
	"errors"
	"sync"
)

type sentinelErr string

func (se sentinelErr) Error() string { return string(se) }

// Closed is a sentinel panic error value that may be raised when
// Generator.Close is called within the provided function.
const Closed sentinelErr = "generator closed"

// A Generator (often called a resumable function or coroutine) is a lazily
// executed function that can be suspended during execution, yielding values
// each time Next is called. When a generator is created, no code is executed
// until Next is called and once a yield is encountered, further execution is
// halted until a subsequent Next.
//
// To ensure underlying resources (mainly goroutines) are cleaned up, Close
// should be deferred after creation.
type Generator[E any] struct {
	cont chan struct{}
	out  chan E

	closeOnce sync.Once
	close     chan struct{}
}

// New creates a Generator powered by the provided fn. The function should call
// yield with the values it wishes to emit, and return if/when the generator is
// complete.
func New[E any](fn func(yield func(E))) *Generator[E] {
	gen := &Generator[E]{
		cont:  make(chan struct{}),
		out:   make(chan E),
		close: make(chan struct{}),
	}

	go gen.run(fn)

	return gen
}

// Next resumes execution of the Generator and returns the next value. If more
// is false, the generator has completed and the value will be the zero-value
// of E.
func (g *Generator[E]) Next() (value E, more bool) {
	select {
	case g.cont <- struct{}{}:
	case <-g.close:
	}

	value, more = <-g.out
	return value, more
}

// Close terminates the Generator, resulting in all subsequent calls to Next to
// return the zero value of E and false. Close can be called multiple times
// safely and should be deferred to clean up resources after the Generator is
// no longer used. This may raise a Closed panic that is handled and
// discarded by the internals of the Generator, allowing any defer statements
// to execute in the provided function.
func (g *Generator[E]) Close() {
	g.closeOnce.Do(func() { close(g.close) })
}

func (g *Generator[E]) wait() {
	select {
	case <-g.cont:
	case <-g.close:
		panic(Closed)
	}
}

func (g *Generator[E]) yield(value E) {
	g.out <- value
	g.wait()
}

func (g *Generator[E]) run(fn func(func(E))) {
	defer func() {
		if rec := recover(); rec != nil {
			if err, ok := rec.(error); !ok || !errors.Is(err, Closed) {
				panic(rec)
			}
		}
	}()

	defer g.Close()
	defer close(g.out)

	g.wait()
	fn(g.yield)
}
