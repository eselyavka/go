package p1254

func dfs(x, y int, grid [][]int) {
	if 0 <= x && x < len(grid) && 0 <= y && y < len(grid[0]) && grid[x][y] == 0 {
		grid[x][y] = 1
		dfs(x+1, y, grid)
		dfs(x-1, y, grid)
		dfs(x, y+1, grid)
		dfs(x, y-1, grid)
	}
}
func closedIsland(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if r == 0 || c == 0 || r == n-1 || c == m-1 {
				dfs(r, c, grid)
			}
		}
	}

	ans := 0
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if grid[r][c] == 0 {
				dfs(r, c, grid)
				ans++
			}
		}
	}
	return ans
}
