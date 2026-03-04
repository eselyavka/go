package solutions

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}

	if isPalindrome(s, 0, len(s)-1) {
		return s
	}

	max := string(s[0])
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j > i; j-- {
			if s[i] == s[j] {
				if isPalindrome(s, i, j) && len(max) < j-i+1 {
					max = s[i : j+1]
				}
			}
		}
	}

	return max
}

func longestPalindromeDP(s string) string {
	n := len(s)
	ans := ""
	dp := make([]bool, n, n)

	for j := 0; j < n; j++ {
		for i := 0; i < j+1; i++ {
			if i == j {
				dp[i] = true
			} else if j == i+1 {
				dp[i] = s[i] == s[j]
			} else {
				flag := dp[i+1] && s[i] == s[j]
				dp[i] = flag
			}
			if dp[i] && j-i+1 > len(ans) {
				ans = s[i : j+1]
			}
		}
	}

	return ans
}
