package twosum_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ashep/pdsa/arrays/twosum"
)

func TestTwoSumV1(t *testing.T) {
	testTwoSum(t, twosum.V1)
}

func TestTwoSumV2(t *testing.T) {
	testTwoSum(t, twosum.V2)
}

func testTwoSum(t *testing.T, fn func(nums []int, target int) []int) {
	t.Run("OneSolution", func(tt *testing.T) {
		require.Equal(tt, []int{1, 2}, fn([]int{3, 2, 4}, 6))
	})

	t.Run("OneSolutionSameNumbers", func(tt *testing.T) {
		require.Equal(tt, []int{0, 1}, fn([]int{3, 3}, 6))
	})

	t.Run("TwoSolutions", func(tt *testing.T) {
		require.Equal(tt, []int{0, 3}, fn([]int{1, 2, 3, 4}, 5))
	})

	t.Run("NoSolutions", func(tt *testing.T) {
		require.Equal(tt, []int(nil), fn([]int{1, 2, 3, 4}, 8))
	})
}

func makeArr(n int) []int {
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = i
	}

	return r
}
