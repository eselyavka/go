package solutions

import "strings"

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		slice := make([]int, n+1)
		dp[i] = slice
	}

	for _, str := range strs {
		z := strings.Count(str, "0")
		o := strings.Count(str, "1")

		if z > m || o > n {
			continue
		}

		for i := m; i >= z; i-- {
			for j := n; j >= o; j-- {
				dp[i][j] = MaxInts([]int{dp[i][j], dp[i-z][j-o] + 1})
			}
		}
	}

	return dp[m][n]
}
