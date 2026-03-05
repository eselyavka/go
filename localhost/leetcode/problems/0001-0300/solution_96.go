package solutions

func numTrees(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res = res * (n + i) / i
	}
	res = res / (n + 1)

	return res
}
