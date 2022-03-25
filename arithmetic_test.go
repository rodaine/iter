package iter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		i := Empty[int]()
		assert.Zero(t, Sum(i))
	})

	t.Run("values", func(t *testing.T) {
		t.Parallel()

		i := FromItems(1, 2, 3)
		assert.Equal(t, 6, Sum(i))
	})
}

func ExampleSum() {
	i := FromItems(1, 2, 3)
	fmt.Println(Sum(i)) // Output: 6
}

func TestProduct(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		i := Empty[int]()
		assert.Zero(t, Product(i))
	})

	t.Run("values", func(t *testing.T) {
		t.Parallel()

		i := FromItems(-1, 2, 3)
		assert.Equal(t, -6, Product(i))
	})
}

func ExampleProduct() {
	i := FromItems(-1, -2, -3)
	fmt.Println(Product(i)) // Output: -6
}
