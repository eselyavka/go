package solutions

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1, n+1)

	for i := 2; i < n+1; i++ {
		one_step := cost[i-1] + dp[i-1]
		two_steps := cost[i-2] + dp[i-2]

		dp[i] = MinInts([]int{one_step, two_steps})
	}

	return dp[n]
}
