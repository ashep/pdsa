package arraysortedmerge_test

import (
	"testing"

	"github.com/ashep/pdsa/misc/arraysortedmerge"
)

// This var kept global intentionally to avoid compiler optimizations
var result []int

func BenchmarkV1(b *testing.B) {
	benchmark(b, arraysortedmerge.V1)
}

func BenchmarkV2(b *testing.B) {
	benchmark(b, arraysortedmerge.V2)
}

func benchmark(bb *testing.B, fn func([]int, []int) []int) {
	var a1Of1K = makeArr(1000)
	var a2Of1K = makeArr(1000)
	var a1Of100K = makeArr(100000)
	var a2Of100K = makeArr(1000000)
	var a1Of1M = makeArr(1000000)
	var a2Of1M = makeArr(1000000)
	var a1Of100M = makeArr(100000000)
	var a2Of100M = makeArr(100000000)

	bb.Run("1K", func(b *testing.B) {
		var r []int

		for n := 0; n < b.N; n++ {
			r = fn(a1Of1K, a2Of1K)
		}

		result = r
	})

	bb.Run("100K", func(b *testing.B) {
		var r []int

		for n := 0; n < b.N; n++ {
			r = fn(a1Of100K, a2Of100K)
		}

		result = r
	})

	bb.Run("1M", func(b *testing.B) {
		var r []int

		for n := 0; n < b.N; n++ {
			r = fn(a1Of1M, a2Of1M)
		}

		result = r
	})

	bb.Run("100M", func(b *testing.B) {
		var r []int

		for n := 0; n < b.N; n++ {
			r = fn(a1Of100M, a2Of100M)
		}

		result = r
	})
}

func makeArr(n int) []int {
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = i
	}

	return r
}
