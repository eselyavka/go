package solutions

func countSquares(matrix [][]int) int {
	row := len(matrix)
	col := len(matrix[0])

	dp := make([][]int, row)
	ans := 0
	for i := range matrix {
		dp[i] = make([]int, col)
		copy(dp[i], matrix[i])
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] == 1 {
				dp[i][j] = 1 + MinInts([]int{dp[i-1][j], dp[i][j-1], dp[i-1][j-1]})
				ans += dp[i][j]
			}
		}
	}

	return ans
}
