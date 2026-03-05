package p1413

import "localhost/leetcode/util"

func minStartValue(nums []int) int {
	ans := util.MinInt

	sum_so_far := 0

	for _, num := range nums {
		sum_so_far += num
		if sum_so_far <= 0 {
			x := -1*sum_so_far + 1
			ans = util.MaxInts([]int{ans, x})
		}
	}

	if ans == util.MinInt {
		return 1
	}

	return ans
}
