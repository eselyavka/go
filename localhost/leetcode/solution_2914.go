package solutions

func minChanges(s string) int {
	n := len(s)
	changes := 0

	for i := 1; i < n; i += 2 {
		if s[i] != s[i-1] {
			changes++
		}
	}

	return changes
}
