package iter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountUp(t *testing.T) {
	t.Parallel()

	i := CountUp(123)

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

	i := CountDown(123)

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

		i := Range(0, 3)

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

		i := Range(3, 0)

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
