package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterator_StepBy(t *testing.T) {
	t.Parallel()

	i := FromItems(0, 1, 2, 3, 4, 5).StepBy(2)

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

	i = FromItems(0, 1, 2).StepBy(1)

	next, ok = i.Next()
	assert.Equal(t, 0, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, 1, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, 2, next)
	assert.True(t, ok)
}
