package iter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromChan(t *testing.T) {
	t.Parallel()

	c := make(chan rune, 3)

	i := FromChan(c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	close(c)

	out, ok := i.Next()
	assert.Equal(t, 'a', out)
	assert.True(t, ok)

	out, ok = i.Next()
	assert.Equal(t, 'b', out)
	assert.True(t, ok)

	out, ok = i.Next()
	assert.Equal(t, 'c', out)
	assert.True(t, ok)

	out, ok = i.Next()
	assert.Zero(t, out)
	assert.False(t, ok)
}

func ExampleFromChan() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	chars := FromChan(ch).ToSlice()
	fmt.Println(chars) // Output: [1 2 3]
}

func TestIterator_ToChan(t *testing.T) {
	t.Parallel()

	t.Run("exhaust", func(t *testing.T) {
		t.Parallel()

		ch, cancel := FromItems(1, 2, 3).ToChan()
		defer cancel()

		assert.Equal(t, 1, <-ch)
		assert.Equal(t, 2, <-ch)
		assert.Equal(t, 3, <-ch)

		out, ok := <-ch
		assert.Zero(t, out)
		assert.False(t, ok)
	})

	t.Run("canceled", func(t *testing.T) {
		t.Parallel()

		ch, cancel := FromItems(1, 2, 3).ToChan()

		assert.Equal(t, 1, <-ch)
		cancel()

		out, ok := <-ch
		assert.Zero(t, out)
		assert.False(t, ok)
	})
}

func ExampleIterator_ToChan() {
	ch, cancel := FromItems(1, 2, 3).ToChan()
	defer cancel()

	fmt.Print(<-ch)
	fmt.Print(<-ch)
	fmt.Println(<-ch) // Output: 123
}
