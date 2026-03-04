package solutions

func countSubstrings(s string) int {
	n := len(s)
	dp := make([]bool, n, n)
	ans := 0
	for j := 0; j < n; j++ {
		for i := 0; i < j+1; i++ {
			if j == i {
				dp[i] = true
			} else if j == i+1 {
				dp[i] = s[i] == s[j]
			} else {
				dp[i] = dp[i+1] && s[i] == s[j]
			}
			if dp[i] {
				ans++
			}
		}
	}
	return ans
}
