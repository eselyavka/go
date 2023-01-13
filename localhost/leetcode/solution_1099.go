package solutions

import "sort"

func twoSumLessThanK(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)

	ans := -1
	left := 0
	right := n - 1

	for left < right {
		sum := nums[left] + nums[right]
		if sum < k {
			ans = MaxInts([]int{sum, ans})
			left++
		} else {
			right--
		}
	}
	return ans
}
