package solutions

func firstCompleteIndex(arr []int, mat [][]int) int {
	m := len(mat)
	n := len(mat[0])

	rowFreq := make(map[int]int)
	for i := 0; i < m; i++ {
		rowFreq[i] = 0
	}
	colFreq := make(map[int]int)
	for j := 0; j < n; j++ {
		colFreq[n] = 0
	}
	vals := make(map[int]tupleInt)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			vals[mat[i][j]] = tupleInt{row: i, col: j}
		}
	}

	for idx, item := range arr {
		t := vals[item]
		rowFreq[t.row]++
		colFreq[t.col]++
		if rowFreq[t.row] == n || colFreq[t.col] == m {
			return idx
		}
	}

	return -1
}
