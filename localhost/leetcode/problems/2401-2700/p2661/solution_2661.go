package p2661

import "localhost/leetcode/util"

func firstCompleteIndex(arr []int, mat [][]int) int {
	m := len(mat)
	n := len(mat[0])

	rowFreq := make(map[int]int)
	for i := 0; i < m; i++ {
		rowFreq[i] = 0
	}
	colFreq := make(map[int]int)
	for j := 0; j < n; j++ {
		colFreq[j] = 0
	}
	vals := make(map[int]util.TupleInt)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			vals[mat[i][j]] = util.TupleInt{Row: i, Col: j}
		}
	}

	for idx, item := range arr {
		t := vals[item]
		rowFreq[t.Row]++
		colFreq[t.Col]++
		if rowFreq[t.Row] == n || colFreq[t.Col] == m {
			return idx
		}
	}

	return -1
}
