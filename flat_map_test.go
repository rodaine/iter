package iter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatMap(t *testing.T) {
	t.Parallel()

	i := FlatMap(FromItems("foo", "bar"), func(s string) Iterator[rune] {
		return FromSlice([]rune(s))
	})

	out := i.ToSlice()
	assert.Equal(t, []rune{'f', 'o', 'o', 'b', 'a', 'r'}, out)
}

func ExampleFlatMap() {
	sl := FlatMap(
		FromItems(1, 2, 3),
		func(n int) Iterator[int] {
			return RangeBy(0, n, 1)
		}).ToSlice()

	fmt.Println(sl) // Output: [0 0 1 0 1 2]
}
