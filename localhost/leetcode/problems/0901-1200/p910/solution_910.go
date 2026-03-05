package p910

import (
	"localhost/leetcode/util"
	"sort"
)

func smallestRangeII(nums []int, k int) int {
	sort.Ints(nums)

	n := len(nums)

	ans := nums[n-1] - nums[0]

	left := nums[0] + k
	right := nums[n-1] - k

	for i := 0; i < n-1; i++ {
		max := util.MaxInts([]int{nums[i] + k, right})
		min := util.MinInts([]int{nums[i+1] - k, left})

		ans = util.MinInts([]int{ans, max - min})
	}

	return ans
}
