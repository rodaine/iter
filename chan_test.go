package iter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromChan(t *testing.T) {
	t.Parallel()

	c := make(chan rune, 3)

	i := FromChan(c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	close(c)

	out, ok := i.Next()
	assert.Equal(t, 'a', out)
	assert.True(t, ok)

	out, ok = i.Next()
	assert.Equal(t, 'b', out)
	assert.True(t, ok)

	out, ok = i.Next()
	assert.Equal(t, 'c', out)
	assert.True(t, ok)

	out, ok = i.Next()
	assert.Zero(t, out)
	assert.False(t, ok)
}
