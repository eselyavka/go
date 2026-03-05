package solutions

func mincostTickets(days []int, costs []int) int {
	maxDay := MaxInts(days) + 1
	dp := make([]int, maxDay)

	daysSet := make(map[int]struct{})
	for _, day := range days {
		daysSet[day] = struct{}{}
	}

	for i := 1; i < maxDay; i++ {
		if _, ok := daysSet[i]; ok {

			dp[i] = MinInts(
				[]int{dp[i-1] + costs[0],
					dp[MaxInts([]int{0, i - 7})] + costs[1],
					dp[MaxInts([]int{0, i - 30})] + costs[2]})
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[maxDay-1]
}
