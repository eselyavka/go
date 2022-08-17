package solutions

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
			arr := Reverse(d[i])
			res = append(res, arr...)
		} else {
			res = append(res, d[i]...)
		}
	}

	return res
}
