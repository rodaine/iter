package iter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterator_Chain(t *testing.T) {
	t.Parallel()

	i := FromItems(0, 1).Chain(FromItems(2, 3))

	next, ok := i.Next()
	assert.Equal(t, 0, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, 1, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, 2, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, 3, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Zero(t, next)
	assert.False(t, ok)

	ci := i.Chain()
	assert.Equal(t, i, ci)
}

func ExampleIterator_Chain() {
	sl := FromItems(0, 1).
		Chain(FromItems(2, 3)).
		ToSlice()
	fmt.Println(sl) // Output: [0 1 2 3]
}
