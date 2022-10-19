package twosum_test

import (
	"testing"

	"github.com/ashep/pdsa/arrays/twosum"
)

// This var kept global intentionally to avoid compiler optimizations
var result []int

func BenchmarkV1(b *testing.B) {
	benchmark(b, twosum.V1)
}

func BenchmarkV2(b *testing.B) {
	benchmark(b, twosum.V2)
}

func benchmark(bb *testing.B, fn func(nums []int, target int) []int) {
	var aOf1K = makeArr(1000)
	var aOf100K = makeArr(100000)

	bb.Run("1K", func(b *testing.B) {
		var r []int

		for n := 0; n < b.N; n++ {
			r = fn(aOf1K, 0)
		}

		result = r
	})

	bb.Run("100K", func(b *testing.B) {
		var r []int

		for n := 0; n < b.N; n++ {
			r = fn(aOf100K, 0)
		}

		result = r
	})
}
