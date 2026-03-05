package solutions

func slidingPuzzle(board [][]int) int {
	m := len(board)
	n := len(board[0])

	findZero := func(matrix [][]int) tupleInt {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if matrix[i][j] == 0 {
					return tupleInt{i, j}
				}
			}
		}
		return tupleInt{-1, -1}
	}

	valid := func(matrix [][]int) bool {
		return intSliceEqual(matrix[0], []int{1, 2, 3}) && intSliceEqual(matrix[1], []int{4, 5, 0})
	}

	if valid(board) {
		return 0
	}

	queue := []tuple5{{0, board}}
	seen := make(map[string]struct{})
	seen[twoDSliceToString(board)] = struct{}{}

	for len(queue) > 0 {
		cnt := len(queue)
		for cnt > 0 {
			item := queue[0]
			queue = queue[1:]
			if valid(item.board) {
				return item.steps
			}
			currIdx := findZero(item.board)
			for _, d := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
				dx := currIdx.row + d[0]
				dy := currIdx.col + d[1]
				if 0 <= dx && dx < m && 0 <= dy && dy < n {
					newItem := copy2DArray(item.board)
					tmp := newItem[currIdx.row][currIdx.col]
					newItem[currIdx.row][currIdx.col] = newItem[dx][dy]
					newItem[dx][dy] = tmp
					newItemTuple := twoDSliceToString(newItem)
					if _, ok := seen[newItemTuple]; !ok {
						seen[newItemTuple] = struct{}{}
						queue = append(queue, tuple5{
							steps: item.steps + 1,
							board: newItem,
						})
					}
				}
			}
			cnt--
		}
	}
	return -1
}
