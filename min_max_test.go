package iter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {
	t.Parallel()

	m, ok := Min(FromItems(1, 2, 3, 4))
	assert.Equal(t, 1, m)
	assert.True(t, ok)

	m, ok = Min(FromItems[int]())
	assert.False(t, ok)
}

func TestMax(t *testing.T) {
	t.Parallel()

	m, ok := Max(FromItems(1, 2, 3, 4))
	assert.Equal(t, 4, m)
	assert.True(t, ok)

	m, ok = Max(FromItems[int]())
	assert.False(t, ok)
}
