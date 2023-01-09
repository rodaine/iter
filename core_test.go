package iter

import (
	"fmt"
)

func ExampleFromCoreFunc() {
	a, b := 1, 1

	fib := FromCoreFunc(func() (int, bool) {
		out := a
		a, b = b, a+b
		return out, true
	})

	fmt.Println(fib.Take(5).ToSlice()) // Output: [1 1 2 3 5]
}
