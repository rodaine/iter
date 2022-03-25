package iter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	t.Parallel()

	m, ok := Min(FromItems(4, 3, 1, 2))
	assert.Equal(t, 1, m)
	assert.True(t, ok)

	m, ok = Min(FromItems[int]())
	assert.False(t, ok)
}

func TestMinBy(t *testing.T) {
	t.Parallel()

	now := time.Now()

	m, ok := FromItems(now, now.Add(time.Minute), now.Add(-time.Minute)).
		MinBy(func(a, b time.Time) bool {
			return a.Before(b)
		})
	assert.True(t, ok)
	assert.True(t, now.Add(-time.Minute).Equal(m))
}

func TestMinByKey(t *testing.T) {
	t.Parallel()

	now := time.Now()

	m, ok := MinByKey(
		FromItems(now, now.Add(time.Minute), now.Add(-time.Minute)),
		time.Time.Unix)

	assert.True(t, ok)
	assert.True(t, now.Add(-time.Minute).Equal(m))
}

func TestMax(t *testing.T) {
	t.Parallel()

	m, ok := Max(FromItems(1, 3, 2, 4))
	assert.Equal(t, 4, m)
	assert.True(t, ok)

	m, ok = Max(FromItems[int]())
	assert.False(t, ok)
}

func TestMaxBy(t *testing.T) {
	t.Parallel()

	now := time.Now()

	m, ok := FromItems(now, now.Add(time.Minute), now.Add(-time.Minute)).
		MaxBy(func(a, b time.Time) bool {
			return a.Before(b)
		})
	assert.True(t, ok)
	assert.True(t, now.Add(time.Minute).Equal(m))
}

func TestMaxByKey(t *testing.T) {
	t.Parallel()

	now := time.Now()

	m, ok := MaxByKey(
		FromItems(now, now.Add(time.Minute), now.Add(-time.Minute)),
		time.Time.Unix)

	assert.True(t, ok)
	assert.True(t, now.Add(time.Minute).Equal(m))
}
