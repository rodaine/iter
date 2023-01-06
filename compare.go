package iter

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Equality is an enum type describing the (in)equality of two values
type Equality int8

const (
	// LessThan indicates that the left-hand value is less than the right-hand
	// value
	LessThan Equality = -1

	// Equal indicates that the left-hand value is equal to the right-hand
	// value
	Equal Equality = 0

	// GreaterThan indicates that the left-hand value is greater than the
	// right-hand value
	GreaterThan Equality = 1
)

// String satisfies the fmt.Stringer interface
func (e Equality) String() string {
	switch e {
	case LessThan:
		return "LessThan"
	case Equal:
		return "Equal"
	case GreaterThan:
		return "GreaterThan"
	default:
		return fmt.Sprintf("Equality(%d)", e)
	}
}

func compare[O constraints.Ordered](a, b O) Equality {
	switch {
	case a == b:
		return Equal
	case a < b:
		return LessThan
	default: // case a > b:
		return GreaterThan
	}
}

// CompareFunc is a function that compares two values, returning their relative
// Equality value.
type CompareFunc[T any] func(a, b T) Equality

// Compare lexicographically compares two iterators element-by-element,
// returning their relative Equality value. For element types that do not
// satisfy constraints.Ordered, Iterator.Compare or CompareByKey can
// be used instead.
func Compare[O constraints.Ordered](a, b Iterator[O]) Equality {
	return a.Compare(b, compare[O])
}

// Compare lexicographically compares two iterators element-by-element via
// the provided CompareFunc, returning their relative Equality value. If the
// element type satisfies constraints.Ordered, Compare can be used instead.
func (iter Iterator[E]) Compare(other Iterator[E], fn CompareFunc[E]) Equality {
	for {
		a, aok := iter.Next()
		b, bok := other.Next()

		switch {
		case !aok && !bok:
			return Equal
		case !aok:
			return LessThan
		case !bok:
			return GreaterThan
		}

		if eq := fn(a, b); eq != Equal {
			return eq
		}
	}
}

// CompareByKey lexicographically compares two iterators element-by-element
// based off the Equality of the values returned by the provided KeyFunc.
func CompareByKey[E any, K constraints.Ordered](a, b Iterator[E], fn KeyFunc[E, K]) Equality {
	mfn := MapFunc[E, K](fn)
	ak := Map(a, mfn)
	bk := Map(b, mfn)
	return Compare(ak, bk)
}
