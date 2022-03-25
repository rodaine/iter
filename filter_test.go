package iter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterator_Filter(t *testing.T) {
	t.Parallel()

	i := FromItems(0, 1, 2, 3, 4, 5).
		Filter(func(n int) bool { return n%2 == 0 })

	next, ok := i.Next()
	assert.Equal(t, 0, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, 2, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, 4, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Zero(t, next)
	assert.False(t, ok)
}

func ExampleIterator_Filter() {
	i := FromItems(0, 1, 2, 3, 4, 5).Filter(func(el int) bool {
		return el%2 == 0
	})

	fmt.Println(i.ToSlice()) // Output: [0 2 4]
}
