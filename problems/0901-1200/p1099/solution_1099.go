package p1099

import (
	"github.com/eseliavka/go/util"
	"sort"
)

func twoSumLessThanK(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)

	ans := -1
	left := 0
	right := n - 1

	for left < right {
		sum := nums[left] + nums[right]
		if sum < k {
			ans = util.MaxInts([]int{sum, ans})
			left++
		} else {
			right--
		}
	}
	return ans
}
