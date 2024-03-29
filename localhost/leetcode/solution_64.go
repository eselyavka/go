package solutions

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, 0)

	for i := 0; i < m; i++ {
		arr := make([]int, n, n)
		dp = append(dp, arr)
	}

	dp[0][0] = grid[0][0]

	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = MinInts([]int{dp[i-1][j], dp[i][j-1]}) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}
