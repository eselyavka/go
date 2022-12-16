package solutions

func bfs1631(mid int, heights [][]int) bool {
	row := len(heights)
	col := len(heights[0])
	adj := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	visited := make(map[tupleInt]bool)
	start := tupleInt{row: 0, col: 0}
	visited[start] = true
	queue := make([]tupleInt, 1)
	queue[0] = start

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.row == row-1 && node.col == col-1 {
			return true
		}

		x := node.row
		y := node.col

		visited[node] = true

		for _, item := range adj {
			dx := item[0]
			dy := item[1]
			new_x := dx + x
			new_y := dy + y

			if !(0 <= new_x && new_x < row) || !(0 <= new_y && new_y < col) {
				continue
			}

			new_node := tupleInt{row: new_x, col: new_y}

			if _, ok := visited[new_node]; ok {
				continue
			}

			curr_diff := Abs(heights[new_x][new_y] - heights[x][y])

			if curr_diff <= mid {
				visited[new_node] = true
				queue = append(queue, new_node)
			}
		}
	}

	return false
}
func minimumEffortPath(heights [][]int) int {
	left := 0
	right := 10000000

	for left < right {
		mid := (left + right) / 2
		if bfs1631(mid, heights) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
