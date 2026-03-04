package solutions

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)

	dp := make([][]int, 0)

	for i := 0; i <= m; i++ {
		tmp := make([]int, n+1, n+1)
		dp = append(dp, tmp)
	}

	for j := n - 1; j > -1; j-- {
		for i := m - 1; i > -1; i-- {
			if text2[j] == text1[i] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = MaxInts([]int{dp[i+1][j], dp[i][j+1]})
			}
		}
	}

	return dp[0][0]
}
