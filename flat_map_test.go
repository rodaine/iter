package iter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlatMap(t *testing.T) {
	t.Parallel()

	i := FlatMap(
		FromItems("foo", "bar"),
		func(s string) Iterator[rune] {
			return FromSlice([]rune(s))
		})

	out := i.ToSlice()
	assert.Equal(t, []rune{'f', 'o', 'o', 'b', 'a', 'r'}, out)
}
