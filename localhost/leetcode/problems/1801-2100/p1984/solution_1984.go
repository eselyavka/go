package p1984

import (
	"localhost/leetcode/util"
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	n := len(nums)

	if n == 1 || k == 1 {
		return 0
	}

	i := 0
	ans := math.MaxInt
	sort.Ints(nums)
	for i <= n-k {
		ans = util.MinInts([]int{ans, nums[i+k-1] - nums[i]})
		i++
	}

	return ans
}
