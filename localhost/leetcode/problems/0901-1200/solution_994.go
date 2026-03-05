package solutions

func could_be_infected(x, y int, grid [][]int) bool {
	if 0 <= x && x < len(grid) && 0 <= y && y < len(grid[0]) {
		if grid[x][y] == 1 {
			return true
		}
	}
	return false
}

func contaminating_process(ts int, grid *[][]int) bool {
	is_infected := false
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for i := 0; i < len(*grid); i++ {
		for j := 0; j < len((*grid)[0]); j++ {
			if (*grid)[i][j] == ts {
				for _, dir := range dirs {
					x := i + dir[0]
					y := j + dir[1]
					if could_be_infected(x, y, *grid) {
						is_infected = true
						(*grid)[x][y] = ts + 1
					}
				}
			}
		}
	}
	return is_infected
}
func orangesRotting(grid [][]int) int {
	ts := 2

	for contaminating_process(ts, &grid) {
		ts++
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return ts - 2
}
