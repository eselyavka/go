package solutions

import "math"

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max_ending_here := 0
	max_so_far := math.MinInt64
	for _, num := range nums {
		max_ending_here += num
		max_so_far = MaxInts([]int{max_so_far, max_ending_here})
		max_ending_here = MaxInts([]int{max_ending_here, 0})
	}
	return int(max_so_far)
}
