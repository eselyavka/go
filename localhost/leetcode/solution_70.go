package solutions

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	i := 2
	for i < n {
		dp[i] = dp[i-1] + dp[i-2]
		i++
	}
	return dp[n-1]
}
