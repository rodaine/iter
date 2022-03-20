package iter

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	t.Parallel()

	i := Map(
		FromItems(1, 2, 3),
		strconv.Itoa)

	next, ok := i.Next()
	assert.Equal(t, "1", next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, "2", next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, "3", next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Zero(t, next)
	assert.False(t, ok)
}

func TestMapWhile(t *testing.T) {
	t.Parallel()

	i := MapWhile(
		FromItems(-2, -1, 0, 1, 2),
		func(n int) (string, bool) {
			return strconv.Itoa(n), n < 0
		})

	next, ok := i.Next()
	assert.Equal(t, "-2", next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, "-1", next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Zero(t, next)
	assert.False(t, ok)
}