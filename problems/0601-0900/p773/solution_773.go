package p773

import "github.com/eseliavka/go/util"

func slidingPuzzle(board [][]int) int {
	m := len(board)
	n := len(board[0])

	findZero := func(matrix [][]int) util.TupleInt {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if matrix[i][j] == 0 {
					return util.TupleInt{Row: i, Col: j}
				}
			}
		}
		return util.TupleInt{Row: -1, Col: -1}
	}

	valid := func(matrix [][]int) bool {
		return util.IntSliceEqual(matrix[0], []int{1, 2, 3}) && util.IntSliceEqual(matrix[1], []int{4, 5, 0})
	}

	if valid(board) {
		return 0
	}

	queue := []util.Tuple5{{Steps: 0, Board: board}}
	seen := make(map[string]struct{})
	seen[util.TwoDSliceToString(board)] = struct{}{}

	for len(queue) > 0 {
		cnt := len(queue)
		for cnt > 0 {
			item := queue[0]
			queue = queue[1:]
			if valid(item.Board) {
				return item.Steps
			}
			currIdx := findZero(item.Board)
			for _, d := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
				dx := currIdx.Row + d[0]
				dy := currIdx.Col + d[1]
				if 0 <= dx && dx < m && 0 <= dy && dy < n {
					newItem := util.Copy2DArray(item.Board)
					tmp := newItem[currIdx.Row][currIdx.Col]
					newItem[currIdx.Row][currIdx.Col] = newItem[dx][dy]
					newItem[dx][dy] = tmp
					newItemTuple := util.TwoDSliceToString(newItem)
					if _, ok := seen[newItemTuple]; !ok {
						seen[newItemTuple] = struct{}{}
						queue = append(queue, util.Tuple5{
							Steps: item.Steps + 1,
							Board: newItem,
						})
					}
				}
			}
			cnt--
		}
	}
	return -1
}
