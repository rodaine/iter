package iter

import "fmt"

func Example() {
	isEven := func(n int) bool { return n%2 == 0 }

	add := func(a, b int) int { return a + b }

	sum, ok := CountUpBy(0, 3). // 0, 3, 6, 9, 12, ...
					Filter(isEven). // 0, 6, 12, 18, ...
					Take(5).        // 0, 6, 12, 18, 24
					Reduce(add)     // 60

	fmt.Println(sum, ok) // Output: 60 true
}
