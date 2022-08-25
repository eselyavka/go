package solutions

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	i := 0
	j := 1

	seen := make(map[byte]struct{})
	res := 0
	buf := 0
	for i < len(s) {
		if _, ok := seen[s[i]]; ok {
			res = MaxInt([]int{res, buf})
			i = j
			j++
			buf = 0
			seen = make(map[byte]struct{})
		} else {
			buf++
			seen[s[i]] = struct{}{}
			i += 1
		}
	}

	return MaxInt([]int{res, buf})
}
