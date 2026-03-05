package solutions

func canPartition(nums []int) bool {
	totalSum := sum(nums)
	if totalSum%2 != 0 {
		return false
	}
	half := totalSum / 2
	n := len(nums)

	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, half+1)
	}
	dp[0][0] = true

	for i := 1; i <= n; i++ {
		curr := nums[i-1]
		for j := 1; j <= half; j++ {
			if j < curr {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-curr]
			}
		}
	}

	return dp[n][half]
}
