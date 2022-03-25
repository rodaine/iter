package iter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestIterator_TryForEach(t *testing.T) {
	t.Parallel()

	evenProduct := 1
	fn := func(n int) error {
		if n%2 != 0 {
			return fmt.Errorf("%d is not even", n)
		}
		evenProduct *= n
		return nil
	}

	err := FromItems(2, 4, 6).TryForEach(fn)
	assert.NoError(t, err)
	assert.Equal(t, 48, evenProduct)

	err = FromItems(1, 2, 3).TryForEach(fn)
	assert.Error(t, err)
}

func TestIterator_Partition(t *testing.T) {
	t.Parallel()

	even, odd := FromItems(0, 1, 2, 3, 4).
		Partition(func(n int) bool { return n%2 == 0 })

	assert.Equal(t, []int{0, 2, 4}, even)
	assert.Equal(t, []int{1, 3}, odd)
}

func TestIterator_All(t *testing.T) {
	t.Parallel()

	evenFn := func(n int) bool { return n%2 == 0 }

	ok := FromItems(2, 4, 6).All(evenFn)
	assert.True(t, ok)
	ok = FromItems(0, 1, 2).All(evenFn)
	assert.False(t, ok)
}

func TestIterator_Any(t *testing.T) {
	t.Parallel()

	evenFn := func(n int) bool { return n%2 == 0 }

	ok := FromItems(1, 2, 3).Any(evenFn)
	assert.True(t, ok)
	ok = FromItems(1, 3, 5).Any(evenFn)
	assert.False(t, ok)
}

func TestIterator_Find(t *testing.T) {
	t.Parallel()

	evenFn := func(n int) bool { return n%2 == 0 }

	match, ok := FromItems(1, 2, 3).Find(evenFn)
	assert.Equal(t, 2, match)
	assert.True(t, ok)

	match, ok = FromItems(1, 3, 5).Find(evenFn)
	assert.Zero(t, match)
	assert.False(t, ok)
}

func TestIterator_TryFind(t *testing.T) {
	t.Parallel()

	evenPositiveFn := func(n int) (bool, error) {
		if n <= 0 {
			return false, fmt.Errorf("%d must be positive", n)
		}
		return n%2 == 0, nil
	}

	match, ok, err := FromItems(1, 2, 3).TryFind(evenPositiveFn)
	assert.Equal(t, 2, match)
	assert.True(t, ok)
	assert.NoError(t, err)

	match, ok, err = FromItems(1, 3, 5).TryFind(evenPositiveFn)
	assert.Zero(t, match)
	assert.False(t, ok)
	assert.NoError(t, err)

	match, ok, err = FromItems(0, 1, 2).TryFind(evenPositiveFn)
	assert.Zero(t, match)
	assert.False(t, ok)
	assert.Error(t, err)
}

func TestIterator_Position(t *testing.T) {
	t.Parallel()

	evenFn := func(n int) bool { return n%2 == 0 }

	pos, ok := FromItems(1, 2, 3).Position(evenFn)
	assert.Equal(t, uint(1), pos)
	assert.True(t, ok)

	pos, ok = FromItems(1, 3, 5).Position(evenFn)
	assert.Zero(t, pos)
	assert.False(t, ok)
}
