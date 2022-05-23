package solutions

func isPalindrome(s string, l int, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l += 1
		r -= 1
	}
	return true
}

func validPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	for l < r {
		if s[l] == s[r] {
			l += 1
			r -= 1
			continue
		} else {
			if isPalindrome(s, l+1, r) || isPalindrome(s, l, r-1) {
				return true
			}
			return false
		}
	}
	return true
}
