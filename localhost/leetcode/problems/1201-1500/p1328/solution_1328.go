package p1328

import "localhost/leetcode/util"

func breakPalindrome(palindrome string) string {
	n := len(palindrome)
	if n == 1 {
		return ""
	}

	i := 0

	b := []byte(palindrome)
	for i < n {
		if b[i] != []byte("a")[0] {
			b[i] = []byte("a")[0]
			break
		}
		i++
	}

	s := string(b)
	if s == util.ReverseString(s) {
		return palindrome[:n-1] + "b"
	}

	return string(b)
}
