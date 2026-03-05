package p2140

import "localhost/leetcode/util"

func dfs(i int, memo []int, q [][]int) int {
	if i >= len(q) {
		return 0
	}

	if memo[i] != 0 {
		return memo[i]
	}

	points := q[i][0]
	bpower := q[i][1]

	memo[i] = util.MaxInts([]int{dfs(i+1, memo, q), points + dfs(i+1+bpower, memo, q)})

	return memo[i]
}
func mostPoints(questions [][]int) int64 {
	n := len(questions)

	memo := make([]int, n, n)

	ans := dfs(0, memo, questions)

	return int64(ans)
}
