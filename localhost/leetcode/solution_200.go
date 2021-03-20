package solutions

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if seen[i][j] {
				continue
			}
			if dfs(grid, seen, m, n, i, j) > 0 {
				res++
			}
		}
	}
	return res
}

func dfs(grid [][]byte, seen [][]bool, m, n, x, y int) int {
	if x >= m || x < 0 || y >= n || y < 0 {
		return 0
	}

	if seen[x][y] {
		return 0
	}

	seen[x][y] = true

	if string(grid[x][y]) == "0" {
		return 0
	}

	return 1 + dfs(grid, seen, m, n, x-1, y) + dfs(grid, seen, m, n, x+1, y) + dfs(grid, seen, m, n, x, y-1) + dfs(grid, seen, m, n, x, y+1)
}
