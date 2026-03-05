package solutions

import "math"

func minOperations_1658(nums []int, x int) int {
	curr := sum(nums)
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
			ans = MinInts([]int{ans, (n - 1 - i) + left})
		}
	}

	if ans == math.MaxInt {
		return -1
	}

	return ans
}
