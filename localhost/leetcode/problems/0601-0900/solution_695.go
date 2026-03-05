package solutions

func bfs695(i, j int, seen *[][]bool, grid [][]int) int {
	if !(0 <= i && i < len(grid)) || !(0 <= j && j < len(grid[0])) {
		return 0
	}

	if (*seen)[i][j] == true {
		return 0
	}

	(*seen)[i][j] = true

	if grid[i][j] == 0 {
		return 0
	}

	return grid[i][j] + bfs695(i+1, j, seen, grid) + bfs695(i-1, j, seen, grid) + bfs695(i, j+1, seen, grid) + bfs695(i, j-1, seen, grid)
}

func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	ans := 0

	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if seen[i][j] == true {
				continue
			}
			var res = bfs695(i, j, &seen, grid)
			ans = MaxInts([]int{ans, res})
		}
	}

	return ans
}
