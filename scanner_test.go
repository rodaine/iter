package iter

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromStringScanner(t *testing.T) {
	t.Parallel()

	sc := bufio.NewScanner(
		strings.NewReader("foo bar baz"))
	sc.Split(bufio.ScanWords)

	i := FromStringScanner(sc)

	next, ok := i.Next()
	assert.Equal(t, "foo", next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, "bar", next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, "baz", next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Zero(t, next)
	assert.False(t, ok)

	assert.NoError(t, sc.Err())
}

func TestFromBytesScanner(t *testing.T) {
	t.Parallel()

	sc := bufio.NewScanner(
		strings.NewReader("foo bar baz"))
	sc.Split(bufio.ScanWords)

	i := FromBytesScanner(sc)

	next, ok := i.Next()
	assert.Equal(t, []byte("foo"), next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, []byte("bar"), next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, []byte("baz"), next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Zero(t, next)
	assert.False(t, ok)

	assert.NoError(t, sc.Err())
}
