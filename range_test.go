package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountUp(t *testing.T) {
	t.Parallel()

	i := CountUpBy(123, 1)

	next, ok := i.Next()
	assert.Equal(t, 123, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, 124, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, 125, next)
	assert.True(t, ok)
}

func TestCountDown(t *testing.T) {
	t.Parallel()

	i := CountDownBy(123, 1)

	next, ok := i.Next()
	assert.Equal(t, 123, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, 122, next)
	assert.True(t, ok)

	next, ok = i.Next()
	assert.Equal(t, 121, next)
	assert.True(t, ok)
}

func TestRange(t *testing.T) {
	t.Parallel()

	t.Run("up", func(t *testing.T) {
		t.Parallel()

		i := RangeBy(0, 3, 1)

		next, ok := i.Next()
		assert.Equal(t, 0, next)
		assert.True(t, ok)

		next, ok = i.Next()
		assert.Equal(t, 1, next)
		assert.True(t, ok)

		next, ok = i.Next()
		assert.Equal(t, 2, next)
		assert.True(t, ok)

		next, ok = i.Next()
		assert.Zero(t, next)
		assert.False(t, ok)
	})

	t.Run("down", func(t *testing.T) {
		t.Parallel()

		i := RangeBy(3, 0, 1)

		next, ok := i.Next()
		assert.Equal(t, 3, next)
		assert.True(t, ok)

		next, ok = i.Next()
		assert.Equal(t, 2, next)
		assert.True(t, ok)

		next, ok = i.Next()
		assert.Equal(t, 1, next)
		assert.True(t, ok)

		next, ok = i.Next()
		assert.Zero(t, next)
		assert.False(t, ok)
	})

}
