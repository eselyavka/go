package solutions

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	dp := make([][]int, 0)

	for i := 0; i < m; i++ {
		tmp := make([]int, n, n)
		dp = append(dp, tmp)
	}

	dp[0][0] = bool2int(0 == obstacleGrid[0][0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}

			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}

			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}
