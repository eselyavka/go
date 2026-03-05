package p53

import "localhost/leetcode/util"

func maxSubArray(nums []int) int {
	ans := util.MinInt
	acc := 0

	for _, num := range nums {
		acc += num
		ans = util.MaxInts([]int{ans, acc})
		acc = util.MaxInts([]int{acc, 0})
	}
	return ans
}
