package solutions

func dfs_1020(x, y int, grid [][]int) {
	if 0 <= x && x < len(grid) && 0 <= y && y < len(grid[0]) && grid[x][y] == 1 {
		grid[x][y] = 0
		dfs_1020(x+1, y, grid)
		dfs_1020(x-1, y, grid)
		dfs_1020(x, y+1, grid)
		dfs_1020(x, y-1, grid)
	}
}

func numEnclaves(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if r == 0 || c == 0 || r == n-1 || c == m-1 {
				dfs_1020(r, c, grid)
			}
		}
	}

	ans := 0
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if grid[r][c] == 1 {
				ans++
			}
		}
	}

	return ans
}
