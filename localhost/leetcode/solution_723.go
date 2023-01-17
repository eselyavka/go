package solutions

type Location struct {
	i, j int
}

func adjacent_elements(board [][]int) map[Location]struct{} {
	m := len(board)
	n := len(board[0])

	s := make(map[Location]struct{})

	for i := 1; i < m-1; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] != 0 && board[i-1][j] == board[i][j] && board[i][j] == board[i+1][j] {
				s[Location{i: i - 1, j: j}] = struct{}{}
				s[Location{i: i, j: j}] = struct{}{}
				s[Location{i: i + 1, j: j}] = struct{}{}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] != 0 && board[i][j-1] == board[i][j] && board[i][j] == board[i][j+1] {
				s[Location{i: i, j: j - 1}] = struct{}{}
				s[Location{i: i, j: j}] = struct{}{}
				s[Location{i: i, j: j + 1}] = struct{}{}
			}
		}
	}

	return s
}

func candyCrush(board [][]int) [][]int {
	n := len(board)
	m := len(board[0])

	res := adjacent_elements(board)
	for len(res) != 0 {
		for key, _ := range res {
			board[key.i][key.j] = 0
		}

		for j := 0; j < n; j++ {
			i := m - 1
			for i != 0 {
				if i < m && board[i][j] == 0 && board[i-1][j] != 0 {
					buf := board[i][j]
					board[i][j] = board[i-1][j]
					board[i-1][j] = buf
					i++
					continue
				}
				i--
			}
		}
		res = adjacent_elements(board)
	}

	return board
}
