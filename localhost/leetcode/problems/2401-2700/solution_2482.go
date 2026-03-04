package solutions

func onesMinusZeros(grid [][]int) [][]int {
	m := len(grid)
	n := len(grid[0])

	d_row := make(map[int][]int)
	d_col := make(map[int][]int)

	var onesRow int
	var zerosRow int
	var onesCol int
	var zerosCol int

	// row traversal
	for idx, _ := range grid {
		onesRow = sum(grid[idx])
		zerosRow = n - onesRow
		d_row[idx] = []int{onesRow, zerosRow}
	}

	// col traversal
	for j := 0; j < n; j++ {
		onesCol = 0
		for i := 0; i < m; i++ {
			onesCol += grid[i][j]
		}
		zerosCol = m - onesCol
		d_col[j] = []int{onesCol, zerosCol}
	}

	ans := make([][]int, 0)
	var arr []int
	for i := 0; i < m; i++ {
		arr = make([]int, n, n)
		ans = append(ans, arr)
		for j := 0; j < n; j++ {
			ans[i][j] = d_row[i][0] + d_col[j][0] - d_row[i][1] - d_col[j][1]
		}
	}

	return ans
}
