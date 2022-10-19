package twosum

// V1 solves the task in an brute force way, scanning nums using nested loop.
// Time complexity is O(n^2), space complexity is O(1).
func V1(nums []int, target int) []int {
	var r []int

	for firstNumIndex := 0; firstNumIndex < len(nums)-1; firstNumIndex++ {
		secondNum := target - nums[firstNumIndex]
		for k := firstNumIndex + 1; k < len(nums); k++ {
			if nums[k] == secondNum {
				r = append(r, firstNumIndex, k)
				return r
			}
		}
	}

	return r
}

// V2 uses additional hash map to store all nums values and their indices. Since hash map has time complexity of O(n),
// this allows to search for a second number faster.
// Time complexity is O(n*2), space complexity is O(n).
func V2(nums []int, target int) []int {
	var r []int
	numsMap := make(map[int]int)

	for i, v := range nums {
		numsMap[v] = i
	}

	for firstNumIndex := 0; firstNumIndex < len(nums)-1; firstNumIndex++ {
		if secondNumIndex, ok := numsMap[target-nums[firstNumIndex]]; ok {
			r = append(r, firstNumIndex, secondNumIndex)
			break
		}
	}

	return r
}
