//go:build go1.23

package iter

import (
	"iter"
)

func (iter Iterator[E]) Seq() iter.Seq[E] {
	return func(yield func(E) bool) {
		for next, ok := iter.Next(); ok; next, ok = iter.Next() {
			if !yield(next) {
				return
			}
		}
	}
}

func Seq2[A, B any](iter Iterator[Pair[A, B]]) iter.Seq2[A, B] {
	return func(yield func(A, B) bool) {
		for next, ok := iter.Next(); ok; next, ok = iter.Next() {
			if !yield(next.A, next.B) {
				return
			}
		}
	}
}
