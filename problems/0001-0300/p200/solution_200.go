package p200

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}

	var dfs200 func(x, y int) int
	dfs200 = func(x, y int) int {
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

		return 1 + dfs200(x-1, y) + dfs200(x+1, y) + dfs200(x, y-1) + dfs200(x, y+1)
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if seen[i][j] {
				continue
			}
			if dfs200(i, j) > 0 {
				res++
			}
		}
	}
	return res
}
