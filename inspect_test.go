package iter

import "fmt"

func ExampleIterator_Inspect() {
	i := FromItems(0, 1, 2, 3, 4).
		Inspect(func(n int) { fmt.Printf("filtering: %d\n", n) }).
		Filter(func(n int) bool { return n%2 == 0 }).
		Inspect(func(n int) { fmt.Printf("even: %d\n", n) })

	fmt.Println(Sum(i))
	// Output:
	// filtering: 0
	// even: 0
	// filtering: 1
	// filtering: 2
	// even: 2
	// filtering: 3
	// filtering: 4
	// even: 4
	// 6
}
