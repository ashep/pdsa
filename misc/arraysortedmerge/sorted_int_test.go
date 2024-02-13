package arraysortedmerge_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ashep/pdsa/misc/arraysortedmerge"
)

func TestSortedIntV1(t *testing.T) {
	testSortedInt(t, arraysortedmerge.V1)
}

func TestSortedIntV2(t *testing.T) {
	testSortedInt(t, arraysortedmerge.V2)
}

func testSortedInt(tt *testing.T, fn func([]int, []int) []int) {
	tt.Run("SameLen", func(t *testing.T) {
		r := fn([]int{1, 2, 3}, []int{4, 5, 6})
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, r)
	})

	tt.Run("FirstEmpty", func(t *testing.T) {
		r := fn([]int{}, []int{4, 5, 6})
		assert.Equal(t, []int{4, 5, 6}, r)
	})

	tt.Run("SecondEmpty", func(t *testing.T) {
		r := fn([]int{1, 2, 3}, []int{})
		assert.Equal(t, []int{1, 2, 3}, r)
	})

	tt.Run("FirstLonger", func(t *testing.T) {
		r := fn([]int{1, 2, 3, 4, 5, 6}, []int{7, 8, 9})
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, r)
	})

	tt.Run("SecondLonger", func(t *testing.T) {
		r := fn([]int{7, 8, 9}, []int{1, 2, 3, 4, 5, 6})
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, r)
	})

	tt.Run("Intersected", func(t *testing.T) {
		r := fn([]int{1, 2, 3}, []int{1, 2, 3})
		assert.Equal(t, []int{1, 1, 2, 2, 3, 3}, r)
	})

	tt.Run("Unsorted", func(t *testing.T) {
		r := fn([]int{2, 3, 1}, []int{6, 4, 5})
		assert.Equal(t, []int{2, 3, 1, 6, 4, 5}, r)
	})
}
