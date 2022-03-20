package iter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterator_Take(t *testing.T) {
	t.Parallel()

	i := FromItems(0, 1, 2, 3, 4).Take(3)

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
	assert.Zero(t, next)
	assert.False(t, ok)
}

func TestIterator_TakeWhile(t *testing.T) {
	t.Parallel()

	i := FromItems(-2, -1, 0, 1, 2).
		TakeWhile(func(n int) bool { return n < 0 })

	next, ok := i.Next()
	assert.Equal(t, -2, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, -1, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Zero(t, next)
	assert.False(t, ok)
}
