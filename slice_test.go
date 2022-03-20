package iter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromSlice(t *testing.T) {
	t.Parallel()

	t.Run("nil slice", func(t *testing.T) {
		t.Parallel()

		var s []string
		i := FromSlice(s)
		out, ok := i.Next()
		assert.Zero(t, out)
		assert.False(t, ok)
	})

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()

		s := []int{} //nolint
		i := FromSlice(s)
		out, ok := i.Next()
		assert.Zero(t, out)
		assert.False(t, ok)
	})

	t.Run("with items", func(t *testing.T) {
		t.Parallel()

		s := []rune{'a', 'b', 'c'}
		i := FromSlice(s)

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
	})
}
