package iter

import "golang.org/x/exp/constraints"

// Number is a constraint describing any Go numerical type that supports
// arithmetic operations.
type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

// Sum computes the sum (addition) of the elements of an iterator. An empty
// iterator returns the zero value of the type.
func Sum[N Number](iter Iterator[N]) (sum N) {
	sum, _ = iter.Reduce(func(a, b N) N { return a + b })
	return sum
}

// Product computes the product (multiplication) of the elements of an
// iterator. An empty iterator returns the zero value of the type.
func Product[N Number](iter Iterator[N]) (product N) {
	prod, _ := iter.Reduce(func(a, b N) N { return a * b })
	return prod
}
