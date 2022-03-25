package iter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFold(t *testing.T) {
	t.Parallel()

	s := Fold(
		FromItems(1, 2, 3), 0,
		func(s int, el int) int {
			return s + el
		})
	assert.Equal(t, 6, s)
}

func ExampleFold() {
	s := Fold(
		FromItems(1, 2, 3), 0,
		func(s int, el int) int { return s + el })
	fmt.Println(s) // Output: 6
}

func TestTryFold(t *testing.T) {
	t.Parallel()

	addEvens := func(s int, el int) (int, error) {
		if el%2 != 0 {
			return 0, fmt.Errorf("%d is not even", el)
		}
		return s + el, nil
	}

	s, err := TryFold(FromItems(0, 2, 4), 0, addEvens)
	assert.NoError(t, err)
	assert.Equal(t, 6, s)

	s, err = TryFold(FromItems(0, 1, 2), 0, addEvens)
	assert.Error(t, err)
	assert.Zero(t, s)
}
