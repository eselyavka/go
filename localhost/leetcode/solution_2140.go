package solutions

func dfs_2140(i int, memo []int, q [][]int) int {
	if i >= len(q) {
		return 0
	}

	if memo[i] != 0 {
		return memo[i]
	}

	points := q[i][0]
	bpower := q[i][1]

	memo[i] = MaxInts([]int{dfs_2140(i+1, memo, q), points + dfs_2140(i+1+bpower, memo, q)})

	return memo[i]
}
func mostPoints(questions [][]int) int64 {
	n := len(questions)

	memo := make([]int, n, n)

	ans := dfs_2140(0, memo, questions)

	return int64(ans)
}
