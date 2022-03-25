package iter

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEquality_String(t *testing.T) {
	t.Parallel()

	tests := map[Equality]string{
		LessThan:    "LessThan",
		Equal:       "Equal",
		GreaterThan: "GreaterThan",
		9:           "Equality(9)",
	}

	for eq, str := range tests {
		assert.Equal(t, str, eq.String())
	}
}

func TestCompare(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		iterA, iterB Iterator[int]
		ex           Equality
	}{
		{
			"both empty",
			Empty[int](),
			Empty[int](),
			Equal,
		},
		{
			"first shorter",
			FromItems(1, 2),
			FromItems(1, 2, 3),
			LessThan,
		},
		{
			"second shorter",
			FromItems(1, 2, 3),
			FromItems(1, 2),
			GreaterThan,
		},
		{
			"equal",
			FromItems(1, 2, 3),
			FromItems(1, 2, 3),
			Equal,
		},
		{
			"less than",
			FromItems(1, 1, 3),
			FromItems(1, 2),
			LessThan,
		},
		{
			"greater than",
			FromItems(1, 3),
			FromItems(1, 2, 3),
			GreaterThan,
		},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.ex, Compare(tc.iterA, tc.iterB))
		})
	}
}

func ExampleCompare() {
	a := FromItems(1, 2, 3)
	b := FromItems(1, 2, 4)
	fmt.Println(Compare(a, b)) // Output: LessThan
}

func TestCompareBy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		iterA, iterB Iterator[int]
		ex           Equality
	}{
		{
			"both empty",
			Empty[int](),
			Empty[int](),
			Equal,
		},
		{
			"first shorter",
			FromItems(1, 2),
			FromItems(1, 2, 3),
			LessThan,
		},
		{
			"second shorter",
			FromItems(1, 2, 3),
			FromItems(1, 2),
			GreaterThan,
		},
		{
			"equal",
			FromItems(1, 2, 3),
			FromItems(1, 2, 3),
			Equal,
		},
		{
			"less than",
			FromItems(1, 1, 3),
			FromItems(1, 2),
			LessThan,
		},
		{
			"greater than",
			FromItems(1, 3),
			FromItems(1, 2, 3),
			GreaterThan,
		},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.ex, tc.iterA.Compare(tc.iterB, func(a, b int) Equality {
				if a == b {
					return Equal
				} else if a < b {
					return LessThan
				}
				return GreaterThan
			}))
		})
	}
}

func TestCompareByKey(t *testing.T) {
	t.Parallel()

	now := time.Now()
	a := FromItems(now, now, now.Add(time.Minute))
	b := FromItems(now, now.Add(-time.Minute), now)

	eq := CompareByKey(a, b, func(time time.Time) int64 { return time.Unix() })
	assert.Equal(t, GreaterThan, eq)
}
