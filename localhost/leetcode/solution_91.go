package solutions

func numDecodings(s string) int {
	if string(s[0]) == "0" {
		return 0
	}

	n := len(s)
	dp := make([]int, n+1, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i < n+1; i++ {
		dp[i] = 0
		if string(s[i-1]) > "0" {
			dp[i] = dp[i-1]
		}

		if string(s[i-2]) == "1" || (string(s[i-2]) == "2" && string(s[i-1]) < "7") {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}
