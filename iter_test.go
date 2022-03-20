package iter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterator_Count(t *testing.T) {
	t.Parallel()

	i := FromItems(1, 2, 3)
	assert.Equal(t, uint(3), i.Count())
	out, ok := i.Next()
	assert.Zero(t, out)
	assert.False(t, ok)
	assert.Equal(t, uint(0), i.Count())
}

func TestIterator_Last(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		i := FromItems[int]()
		last, ok := i.Last()
		assert.Zero(t, last)
		assert.False(t, ok)
	})

	t.Run("single", func(t *testing.T) {
		t.Parallel()

		i := FromItems(1)
		last, ok := i.Last()
		assert.Equal(t, 1, last)
		assert.True(t, ok)
	})

	t.Run("multiple", func(t *testing.T) {
		t.Parallel()

		i := FromItems(1, 2, 3)
		last, ok := i.Last()
		assert.Equal(t, 3, last)
		assert.True(t, ok)
	})
}

func TestIterator_AdvanceBy(t *testing.T) {
	t.Parallel()

	t.Run("exhausted", func(t *testing.T) {
		t.Parallel()

		i := FromItems(1)
		ct, ok := i.AdvanceBy(2)
		assert.Equal(t, uint(1), ct)
		assert.False(t, ok)
	})

	t.Run("enough", func(t *testing.T) {
		t.Parallel()

		i := FromItems(1, 2, 3)
		ct, ok := i.AdvanceBy(2)
		assert.Equal(t, uint(2), ct)
		assert.True(t, ok)
	})
}

func TestIterator_Nth(t *testing.T) {
	t.Parallel()

	t.Run("exhausted", func(t *testing.T) {
		t.Parallel()
		i := FromItems(1)
		nth, ok := i.Nth(3)
		assert.Zero(t, nth)
		assert.False(t, ok)
	})

	t.Run("enough", func(t *testing.T) {
		t.Parallel()
		i := FromItems(1, 2, 3, 4)
		nth, ok := i.Nth(2)
		assert.Equal(t, 3, nth)
		assert.True(t, ok)
	})
}

func TestIterator_ForEach(t *testing.T) {
	t.Parallel()

	product := 1
	FromItems(1, 2, 3, 4).
		ForEach(func(n int) { product *= n })

	assert.Equal(t, 24, product)
}

func TestIterator_Partition(t *testing.T) {
	t.Parallel()

	even, odd := FromItems(0, 1, 2, 3, 4).
		Partition(func(n int) bool { return n%2 == 0 })

	assert.Equal(t, []int{0, 2, 4}, even)
	assert.Equal(t, []int{1, 3}, odd)
}
