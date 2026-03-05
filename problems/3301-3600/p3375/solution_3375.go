package p3375

import "github.com/eseliavka/go/util"

func minOperations(nums []int, k int) int {
	if len(nums) == 0 {
		return -1
	}

	minVal := util.MinInts(nums)

	if minVal < k {
		return -1
	}

	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}

	if set[k] {
		return len(set) - 1
	}

	return len(set)
}
