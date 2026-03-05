package p516

import "localhost/leetcode/util"

func longestPalindromeSubseq(s string) int {
	rev := util.ReverseString(s)

	n := len(s)

	dp := make([][]int, 0)

	for i := 0; i <= n; i++ {
		tmp := make([]int, n+1, n+1)
		dp = append(dp, tmp)
	}

	for i := n - 1; i > -1; i-- {
		for j := n - 1; j > -1; j-- {
			if s[i] == rev[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = util.MaxInts([]int{dp[i+1][j], dp[i][j+1]})
			}
		}
	}

	return dp[0][0]
}
