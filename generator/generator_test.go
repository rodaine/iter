package generator

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerator(t *testing.T) {
	t.Parallel()

	t.Run("no yield", func(t *testing.T) {
		t.Parallel()

		gen := New(func(yield func(int)) { /* noop */ })
		v, ok := gen.Next()
		assert.Zero(t, v)
		assert.False(t, ok)

		v, ok = gen.Next()
		assert.Zero(t, v)
		assert.False(t, ok)
	})

	t.Run("will complete", func(t *testing.T) {
		t.Parallel()

		gen := New(func(yield func(int)) {
			yield(1)
			yield(2)
			yield(3)
		})

		v, ok := gen.Next()
		assert.Equal(t, 1, v)
		assert.True(t, ok)

		v, ok = gen.Next()
		assert.Equal(t, 2, v)
		assert.True(t, ok)

		v, ok = gen.Next()
		assert.Equal(t, 3, v)
		assert.True(t, ok)

		v, ok = gen.Next()
		assert.Zero(t, v)
		assert.False(t, ok)

		v, ok = gen.Next()
		assert.Zero(t, v)
		assert.False(t, ok)
	})

	t.Run("infinite", func(t *testing.T) {
		t.Parallel()

		isClosed := make(chan struct{})
		gen := New(func(yield func(int)) {
			defer close(isClosed)
			ct := 0
			for {
				yield(ct)
				ct++
			}
		})

		v, ok := 0, false
		clock := time.After(10 * time.Millisecond)

	loop:
		for {
			select {
			case <-clock:
				break loop
			default:
				v, ok = gen.Next()
			}
		}

		assert.Greater(t, v, 0)
		assert.True(t, ok)

		gen.Close()

		v, ok = gen.Next()
		assert.Zero(t, v)
		assert.False(t, ok)

		v, ok = gen.Next()
		assert.Zero(t, v)
		assert.False(t, ok)

		select {
		case <-time.After(10 * time.Millisecond):
			t.Fatal("goroutine hasn't been cleaned up")
		case <-isClosed:
			// success!
		}
	})
}
