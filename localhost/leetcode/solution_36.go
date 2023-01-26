package solutions

func isValidSudoku(board [][]byte) bool {
	row := len(board)
	col := row

	square1 := make([]byte, 0)
	square2 := make([]byte, 0)
	square3 := make([]byte, 0)

	for i := 0; i < row; i++ {

		rows := make([]byte, 0)
		cols := make([]byte, 0)

		for j := 0; j < col; j++ {
			if board[i][j] != []byte(".")[0] {
				rows = append(rows, board[i][j])
			}
			if board[j][i] != []byte(".")[0] {
				cols = append(cols, board[j][i])
			}

			if j <= 2 && board[i][j] != []byte(".")[0] {
				square1 = append(square1, board[i][j])
			} else if 2 < j && j <= 5 && board[i][j] != []byte(".")[0] {
				square2 = append(square2, board[i][j])
			} else if 5 < j && j <= 8 && board[i][j] != []byte(".")[0] {
				square3 = append(square3, board[i][j])
			}
		}

		if i == 2 {
			if !arrayUniq(square1) || !arrayUniq(square2) || !arrayUniq(square3) {
				return false
			}
			square1 = make([]byte, 0)
			square2 = make([]byte, 0)
			square3 = make([]byte, 0)
		}

		if i == 5 {
			if !arrayUniq(square1) || !arrayUniq(square2) || !arrayUniq(square3) {
				return false
			}
			square1 = make([]byte, 0)
			square2 = make([]byte, 0)
			square3 = make([]byte, 0)
		}

		if i == 8 {
			if !arrayUniq(square1) || !arrayUniq(square2) || !arrayUniq(square3) {
				return false
			}
			square1 = make([]byte, 0)
			square2 = make([]byte, 0)
			square3 = make([]byte, 0)
		}

		if !arrayUniq(rows) || !arrayUniq(cols) {
			return false
		}
	}

	return true
}
