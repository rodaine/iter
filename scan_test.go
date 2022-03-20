package iter

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestScan(t *testing.T) {
	t.Parallel()

	i := Scan(
		FromItems(1, 2, 3), 1,
		func(st *int, n int) (string, bool) {
			*st *= n
			return strconv.Itoa(-*st), true
		})

	next, ok := i.Next()
	assert.Equal(t, "-1", next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, "-2", next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, "-6", next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Zero(t, next)
	assert.False(t, ok)
}
