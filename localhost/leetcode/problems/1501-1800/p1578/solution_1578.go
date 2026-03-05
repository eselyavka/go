package p1578

import "localhost/leetcode/util"

func minCost(colors string, neededTime []int) int {
	ans := 0
	best := 0

	var prev rune

	for idx, curr := range colors {
		if curr == prev {
			ans += util.MinInts([]int{best, neededTime[idx]})
			best = util.MaxInts([]int{best, neededTime[idx]})
		} else {
			prev = curr
			best = neededTime[idx]
		}
	}

	return ans
}
