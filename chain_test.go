package iter

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
}
