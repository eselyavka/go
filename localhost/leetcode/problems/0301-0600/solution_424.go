package solutions

func characterReplacement(s string, k int) int {
	seen := make(map[rune]int)
	l := 0
	ans := 0

	for r, c := range s {
		seen[c]++
		maxf := -1
		for _, v := range seen {
			if v > maxf {
				maxf = v
			}
		}
		for (r-l+1)-maxf > k {
			seen[rune(s[l])]--
			l++
		}
		ans = MaxInts([]int{ans, r - l + 1})
	}

	return ans
}
