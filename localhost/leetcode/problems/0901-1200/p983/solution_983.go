package p983

import "localhost/leetcode/util"

func mincostTickets(days []int, costs []int) int {
	maxDay := util.MaxInts(days) + 1
	dp := make([]int, maxDay)

	daysSet := make(map[int]struct{})
	for _, day := range days {
		daysSet[day] = struct{}{}
	}

	for i := 1; i < maxDay; i++ {
		if _, ok := daysSet[i]; ok {

			dp[i] = util.MinInts(
				[]int{dp[i-1] + costs[0],
					dp[util.MaxInts([]int{0, i - 7})] + costs[1],
					dp[util.MaxInts([]int{0, i - 30})] + costs[2]})
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[maxDay-1]
}
