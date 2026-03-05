package p1658

import (
	"localhost/leetcode/util"
	"math"
)

func minOperations_1658(nums []int, x int) int {
	curr := util.Sum(nums)
	ans := math.MaxInt
	left := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		curr -= nums[i]

		for curr < x && left <= i {
			curr += nums[left]
			left++
		}

		if curr == x {
			ans = util.MinInts([]int{ans, (n - 1 - i) + left})
		}
	}

	if ans == math.MaxInt {
		return -1
	}

	return ans
}
