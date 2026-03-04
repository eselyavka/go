package solutions

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	mapping := make(map[rune]int)
	for _, c := range t {
		mapping[c]++
	}

	seen := make(map[rune]int)

	l := 0
	cnt := 0
	n := len(t)
	max_int := int(^uint(0) >> 1)

	ans := [3]int{max_int, -1, -1}

	for r, c := range s {
		seen[c]++

		if _, ok := mapping[c]; ok && seen[c] <= mapping[c] {
			cnt++
		}

		for l <= r && cnt == n {
			ch := rune(s[l])
			if r-l+1 < ans[0] {
				ans[0] = r - l + 1
				ans[1] = l
				ans[2] = r
			}
			seen[ch]--
			if _, ok := mapping[ch]; ok && seen[ch] < mapping[ch] {
				cnt--
			}
			l++
		}
	}

	if ans[0] == int(^uint(0)>>1) {
		return ""
	}

	return s[ans[1] : ans[2]+1]
}
