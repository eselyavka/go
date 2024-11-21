package solutions

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	grid := make([][]string, m)
	for i := range grid {
		grid[i] = make([]string, n)
	}

	for _, guard := range guards {
		grid[guard[0]][guard[1]] = "G"
	}

	for _, wall := range walls {
		grid[wall[0]][wall[1]] = "W"
	}

	seen := make(map[tupleInt]struct{})

	for _, guard := range guards {
		x := guard[0]
		y := guard[1]
		dx := x + 1
		for dx < m && grid[dx][y] == "" {
			seen[tupleInt{row: dx, col: y}] = struct{}{}
			dx++
		}
		dx = x - 1
		for dx >= 0 && grid[dx][y] == "" {
			seen[tupleInt{row: dx, col: y}] = struct{}{}
			dx--
		}
		dy := y + 1
		for dy < n && grid[x][dy] == "" {
			seen[tupleInt{row: x, col: dy}] = struct{}{}
			dy++
		}
		dy = y - 1
		for dy >= 0 && grid[x][dy] == "" {
			seen[tupleInt{row: x, col: dy}] = struct{}{}
			dy--
		}
	}

	return m*n - len(seen) - len(walls) - len(guards)
}
