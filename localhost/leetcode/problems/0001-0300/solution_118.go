package solutions

func generate(numRows int) [][]int {
	res := make([][]int, numRows)

	for i := range res {
		res[i] = make([]int, i+1)
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j < len(res[i]); j++ {
			res[i][j] = 1
		}
	}

	for i := 0; i < numRows; i++ {
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}

	return res
}
