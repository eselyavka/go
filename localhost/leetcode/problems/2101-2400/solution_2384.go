package solutions

import "strings"

func largestPalindromic(num string) string {
	n := len(num)
	zeros := make([]rune, n)
	for i := 0; i < n; i++ {
		zeros[i] = []rune("0")[0]
	}
	if num == string(zeros) {
		return "0"
	}

	cnt := make(map[rune]int)
	for _, c := range num {
		cnt[c]++
	}

	ans := make([]rune, n, n)
	l := 0
	r := n - 1
	mid := -1
	for _, c := range []rune("9876543210") {
		if _, ok := cnt[c]; ok {
			occ := cnt[c]
			if occ%2 == 0 {
				for occ != 0 {
					ans[l] = c
					ans[r] = c
					occ -= 2
					l++
					r--
				}
			} else {
				if mid < 0 {
					mid = n / 2
					ans[mid] = c
				}
				occ--
				for occ != 0 {
					ans[l] = c
					ans[r] = c
					occ -= 2
					l++
					r--
				}
			}
		}
	}

	l = 0
	r = n - 1
	i := 0
	for i < n {
		if ans[l] == 48 {
			ans[l] = []rune("|")[0]
			ans[r] = []rune("|")[0]
			l++
			r--
		}
		if ans[i] == 0 {
			ans[i] = []rune("|")[0]
		}
		i++
	}

	return strings.Replace(string(ans), "|", "", -1)
}
