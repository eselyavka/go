package p485

import "localhost/leetcode/util"

func findMaxConsecutiveOnes(nums []int) int {
	runningSum := 0
	prevSum := 0
	for _, n := range nums {
		if n == 1 {
			runningSum += 1
		} else {
			prevSum = util.MaxInts([]int{prevSum, runningSum})
			runningSum = 0
		}
	}

	return util.MaxInts([]int{runningSum, prevSum})
}
