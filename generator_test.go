package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromGeneratorFunc(t *testing.T) {
	t.Parallel()

	s := FromGeneratorFunc(func(yield func(int)) {
		prev, curr := 0, 1
		for {
			yield(curr)
			prev, curr = curr, prev+curr
		}
	}).Take(10).ToSlice()

	assert.Equal(t, []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}, s)
}
