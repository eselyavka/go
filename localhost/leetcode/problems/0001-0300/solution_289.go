package solutions

func isLive(x, y int, board [][]int) bool {
	if x >= 0 && x < len(board) && y >= 0 && y < len(board[0]) {
		return board[x][y] == 1
	}

	return false
}

func gameOfLife(board [][]int) {
	m := len(board)
	n := len(board[0])

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, 1}, {1, 1}, {-1, -1}, {1, -1}}

	changes := make([][]int, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			live := 0
			for _, dir := range dirs {
				newx := i + dir[0]
				newy := j + dir[1]
				if isLive(newx, newy, board) {
					live++
				}
			}

			if live < 2 {
				changes = append(changes, []int{i, j, 0})
			}

			if live > 3 {
				changes = append(changes, []int{i, j, 0})
			}

			if live == 3 && board[i][j] == 0 {
				changes = append(changes, []int{i, j, 1})
			}
		}
	}

	for _, change := range changes {
		x := change[0]
		y := change[1]
		val := change[2]

		board[x][y] = val
	}
}
