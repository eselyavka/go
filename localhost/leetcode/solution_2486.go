package solutions

func appendCharacters(s string, t string) int {
	if s == t {
		return 0
	}

	i := 0
	j := 0
	n := len(s)
	m := len(t)

	ans := m

	for i < n && j < m {
		if s[i] == t[j] {
			i++
			j++
			ans--
		} else {
			i++
		}
	}
	return ans
}
