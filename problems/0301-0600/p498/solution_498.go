package p498

import "github.com/eseliavka/go/util"

func findDiagonalOrder(mat [][]int) []int {
	row := len(mat)
	col := len(mat[0])
	if row == 1 && col == 1 {
		return []int{mat[0][0]}
	}

	d := make(map[int][]int, 0)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			key := i + j
			d[key] = append(d[key], mat[i][j])
		}
	}

	res := make([]int, 0)

	for i := 0; i < len(d); i++ {
		if i%2 == 0 {
			arr := util.Reverse(d[i])
			res = append(res, arr...)
		} else {
			res = append(res, d[i]...)
		}
	}

	return res
}

func findDiagonalOrder2(mat [][]int) []int {
	rows := len(mat)
	cols := len(mat[0])
	res := make([]int, 0, rows*cols)
	row := 0
	col := 0
	direction := 1

	for i := 0; i < rows*cols; i++ {
		res = append(res, mat[row][col])

		if direction == 1 {
			if col == cols-1 {
				row++
				direction = -1
			} else if row == 0 {
				col++
				direction = -1
			} else {
				row--
				col++
			}
		} else {
			if row == rows-1 {
				col++
				direction = 1
			} else if col == 0 {
				row++
				direction = 1
			} else {
				row++
				col--
			}
		}
	}

	return res
}
