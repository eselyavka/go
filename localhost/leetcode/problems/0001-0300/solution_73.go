package solutions

func setZeroes(matrix [][]int) {
	n := len(matrix)
	m := len(matrix[0])

	row := make(map[int]struct{})
	col := make(map[int]struct{})

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				row[i] = struct{}{}
				col[j] = struct{}{}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if _, ok := row[i]; ok {
				matrix[i][j] = 0
			}
			if _, ok := col[j]; ok {
				matrix[i][j] = 0
			}
		}
	}
}
