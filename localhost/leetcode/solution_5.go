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
