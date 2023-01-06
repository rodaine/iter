package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZip(t *testing.T) {
	t.Parallel()

	t.Run("a longer", func(t *testing.T) {
		t.Parallel()

		a := FromItems(0, 1, 2)
		b := FromItems(3, 4)
		i := Zip(a, b)

		next, ok := i.Next()
		assert.Equal(t, Pair[int, int]{0, 3}, next)
		assert.True(t, ok)

		next, ok = i.Next()
		assert.Equal(t, Pair[int, int]{1, 4}, next)
		assert.True(t, ok)

		next, ok = i.Next()
		assert.Zero(t, next)
		assert.False(t, ok)
	})

	t.Run("b longer", func(t *testing.T) {
		t.Parallel()

		a := FromItems(0, 1)
		b := FromItems(2, 3, 4)
		i := Zip(a, b)

		next, ok := i.Next()
		assert.Equal(t, Pair[int, int]{0, 2}, next)
		assert.True(t, ok)

		next, ok = i.Next()
		assert.Equal(t, Pair[int, int]{1, 3}, next)
		assert.True(t, ok)

		next, ok = i.Next()
		assert.Zero(t, next)
		assert.False(t, ok)
	})
}

func TestEnumerate(t *testing.T) {
	t.Parallel()

	i := Enumerate(FromItems('a', 'b', 'c'))

	next, ok := i.Next()
	assert.Equal(t, Pair[int, rune]{0, 'a'}, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, Pair[int, rune]{1, 'b'}, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, Pair[int, rune]{2, 'c'}, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Zero(t, next)
	assert.False(t, ok)
}
