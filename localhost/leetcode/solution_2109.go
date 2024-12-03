package solutions

func addSpaces(s string, spaces []int) string {
	ans := make([]byte, len(s)+len(spaces))

	spaceIdx := 0
	stringIdx := 0

	i := 0
	for i < len(s) {
		if spaceIdx < len(spaces) && i == spaces[spaceIdx] {
			ans[stringIdx] = ' '
			stringIdx++
			spaceIdx++
		}
		ans[stringIdx] = s[i]
		i++
		stringIdx++
	}

	return string(ans)
}
