package p1493

import "localhost/leetcode/util"

func longestSubarray(nums []int) int {
	n := len(nums)
	if util.Sum(nums) == n {
		return n - 1
	}

	ans := 0
	i := 0
	prevCnt := 0
	currCnt := 0

	for i < n {
		if nums[i] == 1 {
			currCnt++
		} else {
			ans = util.MaxInts([]int{ans, prevCnt + currCnt})
			prevCnt = currCnt
			currCnt = 0
		}
		i++
	}

	return util.MaxInts([]int{ans, prevCnt + currCnt})
}
