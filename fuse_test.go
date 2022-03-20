package iter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterator_Fuse(t *testing.T) {
	t.Parallel()

	i := Scan(CountUp(0), 0,
		func(st *int, n int) (int, bool) {
			*st++
			return n, *st%2 == 0
		})

	next, ok := i.Next()
	assert.False(t, ok)

	next, ok = i.Next()
	assert.Equal(t, 1, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.False(t, ok)

	next, ok = i.Next()
	assert.Equal(t, 3, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.False(t, ok)

	i = i.Fuse()

	next, ok = i.Next()
	assert.Equal(t, 5, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.False(t, ok)

	next, ok = i.Next()
	assert.False(t, ok)
}
