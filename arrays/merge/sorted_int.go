// Package merge demonstrates ways of merging two sorted arrays of integers.
package merge

// SortedIntV1 uses slice shifting.
// As benchmark tests show, this version works slightly faster than v2.
func SortedIntV1(a1, a2 []int) []int {
	if len(a1) == 0 {
		return a2
	}

	if len(a2) == 0 {
		return a1
	}

	// Pre-allocate space for result to make append() work faster
	res := make([]int, 0, len(a1)+len(a2))

	for {
		if len(a1)+len(a2) == 0 {
			break
		}

		// No more items in first array
		if len(a1) == 0 {
			res = append(res, a2[0])
			a2 = a2[1:]
			continue
		}

		// No more items in second array
		if len(a2) == 0 {
			res = append(res, a1[0])
			a1 = a1[1:]
			continue
		}

		if a1[0] <= a2[0] {
			res = append(res, a1[0])
			a1 = a1[1:]
		} else {
			res = append(res, a2[0])
			a2 = a2[1:]
		}
	}

	return res
}

// SortedIntV2 uses indices to track each array during looping.
// As benchmark tests show, this version works slightly slower than v1.
func SortedIntV2(a1, a2 []int) []int {
	if len(a1) == 0 {
		return a2
	}

	if len(a2) == 0 {
		return a1
	}

	// Pre-allocate space for result to make append() work faster
	res := make([]int, 0, len(a1)+len(a2))

	i1 := 0
	i2 := 0
	for k := 0; k < len(a1)+len(a2); k++ {
		// No more items in first array
		if i1 == len(a1) {
			res = append(res, a2[i2])
			i2++
			continue
		}

		// No more items in second array
		if i2 == len(a2) {
			res = append(res, a1[i1])
			i1++
			continue
		}

		if a1[i1] <= a2[i2] {
			res = append(res, a1[i1])
			i1++
		} else {
			res = append(res, a2[i2])
			i2++
		}
	}

	return res
}
