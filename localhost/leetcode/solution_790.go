package solutions

func numTilings(n int) int {
	ans := make([]int, 10001, 10001)
	ans[1] = 1
	ans[2] = 2
	ans[3] = 5

	for i := 4; i <= n; i++ {
		ans[i] = 2*ans[i-1] + ans[i-3]
	}

	return ans[n]
}
