package solutions

func dfs_79(x, y, pos int, board [][]byte, path [][]int, word string) bool {
	if pos == len(word) {
		return true
	}

	if 0 <= x && x < len(board) && 0 <= y && y < len(board[0]) && board[x][y] == word[pos] && !in2dIntArray(path, []int{x, y}) {
		path = append(path, []int{x, y})

		res := dfs_79(x+1, y, pos+1, board, path, word) || dfs_79(x-1, y, pos+1, board, path, word) || dfs_79(x, y+1, pos+1, board, path, word) || dfs_79(x, y-1, pos+1, board, path, word)

		path = path[:len(path)-1]

		return res
	}

	return false

}

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])

	path := make([][]int, 0)

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if dfs_79(r, c, 0, board, path, word) {
				return true
			}
		}
	}

	return false
}
