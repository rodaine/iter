package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterator_Skip(t *testing.T) {
	t.Parallel()

	i := FromItems(0, 1, 2, 3, 4).Skip(3)

	next, ok := i.Next()
	assert.Equal(t, 3, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, 4, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Zero(t, next)
	assert.False(t, ok)
}

func TestIterator_SkipWhile(t *testing.T) {
	t.Parallel()

	i := FromItems(-2, -1, 0, 1, -2).
		SkipWhile(func(n int) bool { return n < 0 })

	next, ok := i.Next()
	assert.Equal(t, 0, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, 1, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, -2, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Zero(t, next)
	assert.False(t, ok)
}
