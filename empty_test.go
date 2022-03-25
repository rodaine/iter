package iter

import "fmt"

func ExampleEmpty() {
	i := Empty[int]()
	_, ok := i.Next()
	fmt.Println(ok) // Output: false
}
