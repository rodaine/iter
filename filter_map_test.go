package iter

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterMap(t *testing.T) {
	t.Parallel()

	i := FilterMap(
		FromItems(0, 1, 2, 3, 4, 5),
		func(n int) (string, bool) {
			return strconv.Itoa(n), n%2 == 0
		})

	next, ok := i.Next()
	assert.Equal(t, "0", next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, "2", next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, "4", next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Zero(t, next)
	assert.False(t, ok)
}

func ExampleFilterMap() {
	sl := FilterMap(
		FromItems(0, 1, 2, 12),
		func(n int) (string, bool) {
			return fmt.Sprintf("%x", n), n%2 == 0
		}).ToSlice()

	fmt.Println(sl) // Output: [0 2 c]
}
