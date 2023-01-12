package solutions

func winnerOfGame(colors string) bool {
	n := len(colors)

	if n < 3 {
		return false
	}

	alice := 0
	bob := 0

	for i := 1; i < n-1; i++ {
		if colors[i-1] == 'A' && colors[i] == 'A' && colors[i+1] == 'A' {
			alice++
		}

		if colors[i-1] == 'B' && colors[i] == 'B' && colors[i+1] == 'B' {
			bob++
		}
	}

	return bob < alice
}
