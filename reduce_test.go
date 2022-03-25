package iter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterator_Reduce(t *testing.T) {
	t.Parallel()

	fn := func(a, b int) int { return a + b }

	sum, ok := FromItems(1, 2, 3).Reduce(fn)
	assert.Equal(t, 6, sum)
	assert.True(t, ok)

	_, ok = Empty[int]().Reduce(fn)
	assert.False(t, ok)
}

func TestIterator_TryReduce(t *testing.T) {
	t.Parallel()

	fn := func(a, b int) (int, error) {
		if b%2 != 0 {
			return 0, fmt.Errorf("%d is not even", b)
		}
		return a + b, nil
	}

	sum, ok, err := FromItems(2, 4, 6).TryReduce(fn)
	assert.Equal(t, 12, sum)
	assert.True(t, ok)
	assert.NoError(t, err)

	_, ok, err = Empty[int]().TryReduce(fn)
	assert.False(t, ok)
	assert.NoError(t, err)

	_, _, err = FromItems(1, 2, 3).TryReduce(fn)
	assert.Error(t, err)
}
