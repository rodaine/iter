package iter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterator_Peek(t *testing.T) {
	t.Parallel()

	i := FromItems(0, 1)

	peek, ok := i.Peek()
	assert.Equal(t, 0, peek)
	assert.True(t, ok)

	next, ok := i.Next()
	assert.Equal(t, 0, next)
	assert.True(t, ok)

	peek, ok = i.Peek()
	assert.Equal(t, 1, peek)
	assert.True(t, ok)

	peek, ok = i.Peek()
	assert.Equal(t, 1, peek)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, 1, next)
	assert.True(t, ok)

	peek, ok = i.Peek()
	assert.Zero(t, peek)
	assert.False(t, ok)
}
