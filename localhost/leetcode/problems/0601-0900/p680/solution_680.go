package p680

import "localhost/leetcode/util"

func validPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	for l < r {
		if s[l] == s[r] {
			l += 1
			r -= 1
			continue
		} else {
			if util.IsPalindrome(s, l+1, r) || util.IsPalindrome(s, l, r-1) {
				return true
			}
			return false
		}
	}
	return true
}
