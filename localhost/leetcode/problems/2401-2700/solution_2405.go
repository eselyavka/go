package solutions

func partitionString(s string) int {
	if s == "" {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	n := len(s)

	seen := make(map[string]struct{})

	seen[string(s[0])] = struct{}{}
	ans := 1
	for i := 1; i < n; i++ {
		if _, ok := seen[string(s[i])]; ok {
			seen = make(map[string]struct{})
			ans++
		}
		seen[string(s[i])] = struct{}{}
	}
	return ans
}
