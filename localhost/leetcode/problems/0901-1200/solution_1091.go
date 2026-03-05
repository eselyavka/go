package solutions

import "strconv"

func neighbours(row, col, n, m int, directions, grid [][]int) [][]int {
	neighbours := make([][]int, 0)

	for _, direction := range directions {
		new_row := row + direction[0]
		new_col := col + direction[1]

		if !(0 <= new_row && new_row <= n && 0 <= new_col && new_col <= m) {
			continue
		}

		if grid[new_row][new_col] == 1 {
			continue
		}

		neighbours = append(neighbours, []int{new_row, new_col})
	}
	return neighbours
}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid) - 1
	m := len(grid[0]) - 1

	if grid[0][0] == 1 || grid[n][m] == 1 {
		return -1
	}

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

	queue := [][]int{{0, 0, 1}}
	visited := make(map[string]struct{})
	for len(queue) > 0 {
		row := queue[0][0]
		col := queue[0][1]
		distance := queue[0][2]

		if row == n && col == m {
			return distance
		}

		for _, neighbour := range neighbours(row, col, n, m, directions, grid) {
			key := strconv.Itoa(neighbour[0]) + strconv.Itoa(neighbour[1])
			if _, ok := visited[key]; ok {
				continue
			}
			visited[key] = struct{}{}
			queue = append(queue, []int{neighbour[0], neighbour[1], distance + 1})
		}

		queue = queue[1:]
	}

	return -1
}
