package p1582

func numSpecial(mat [][]int) int {
	rows := len(mat)
	cols := len(mat[0])

	rowCount := make([]int, rows)
	colCount := make([]int, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if mat[i][j] == 1 {
				rowCount[i]++
				colCount[j]++
			}
		}
	}

	special := 0

	for i := 0; i < rows; i++ {
		if rowCount[i] != 1 {
			continue
		}
		for j := 0; j < cols; j++ {
			if mat[i][j] == 1 && colCount[j] == 1 {
				special++
			}
		}
	}

	return special
}
