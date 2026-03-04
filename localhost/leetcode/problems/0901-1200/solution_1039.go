package solutions

func rec_1039(memo [][]int, arr []int, i, j int) int {
	if i+1 == j {
		return 0
	}

	if memo[i][j] != -1 {
		return memo[i][j]
	}

	res := MaxInt

	for k := i + 1; k < j; k++ {
		curr := rec_1039(memo, arr, i, k) + rec_1039(memo, arr, k, j) + arr[i]*arr[j]*arr[k]
		res = MinInts([]int{res, curr})
	}
	memo[i][j] = res

	return res
}

func minScoreTriangulation(values []int) int {
	n := len(values)
	memo := make([][]int, 0)
	for i := 0; i < len(values); i++ {
		memo = append(memo, make([]int, n))
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}

	ans := rec_1039(memo, values, 0, n-1)
	return ans
}
