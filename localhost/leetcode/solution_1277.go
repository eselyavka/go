package solutions

func countSquares(matrix [][]int) int {
	row := len(matrix)
	col := len(matrix[0])

	dp := make([][]int, row+1)
	ans := 0
	for i := 0; i < row+1; i++ {
		dp[i] = make([]int, col+1)
	}

	for i := 1; i < row+1; i++ {
		for j := 1; j < col+1; j++ {
			if matrix[i-1][j-1] == 1 {
				dp[i][j] = 1 + MinInts([]int{dp[i-1][j], dp[i][j-1], dp[i-1][j-1]})
				ans += dp[i][j]
			}
		}
	}

	return ans
}
