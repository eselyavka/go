package solutions

import "sort"

func minIncrementForUnique(nums []int) int {
	n := len(nums)

	ans := 0
	if n == 1 {
		return ans
	}

	sort.Ints(nums)
	incr := 0
	for i := 1; i < n; i++ {
		if nums[i] <= nums[i-1] {
			incr += nums[i-1] + 1 - nums[i]
			nums[i] = nums[i-1] + 1
		}
	}
	return incr
}
