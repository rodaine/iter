//go:build go1.23

package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterator_Seq(t *testing.T) {
	t.Parallel()

	items := FromItems(1, 2, 3)
	for item := range items.Seq() {
		if item == 2 {
			break
		}
	}

	next, ok := items.Next()
	assert.True(t, ok)
	assert.Equal(t, 3, next)
}

func TestIterator_Seq2(t *testing.T) {
	t.Parallel()

	items := Enumerate(FromItems(1, 2, 3))
	for _, n := range Seq2(items) {
		if n == 2 {
			break
		}
	}
	next, ok := items.Next()
	assert.True(t, ok)
	assert.Equal(t, 2, next.A)
	assert.Equal(t, 3, next.B)
}
