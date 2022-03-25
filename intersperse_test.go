package iter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterator_Intersperse(t *testing.T) {
	t.Parallel()

	i := FromItems(0, 1, 2, 3).Intersperse(9)

	next, ok := i.Next()
	assert.Equal(t, 0, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, 9, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, 1, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, 9, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, 2, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, 9, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, 3, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Zero(t, next)
	assert.False(t, ok)
}

func ExampleIterator_Intersperse() {
	sl := FromItems(0, 1, 2).Intersperse(3).ToSlice()
	fmt.Println(sl) // Output: [0 3 1 3 2]
}
