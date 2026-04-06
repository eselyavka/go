package p3070

func countSubmatrices(grid [][]int, k int) int {
	rows := len(grid)
	cols := len(grid[0])

	prefix := make([][]int, rows)
	for row := 0; row < rows; row++ {
		prefix[row] = make([]int, cols)
	}

	ans := 0

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			prefix[row][col] = grid[row][col]

			if row > 0 {
				prefix[row][col] += prefix[row-1][col]
			}
			if col > 0 {
				prefix[row][col] += prefix[row][col-1]
			}
			if row > 0 && col > 0 {
				prefix[row][col] -= prefix[row-1][col-1]
			}

			if prefix[row][col] <= k {
				ans++
			}
		}
	}

	return ans
}
