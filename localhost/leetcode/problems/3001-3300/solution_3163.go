package solutions

func compressedString(word string) string {
	comm := make([]rune, 0)
	i := 0

	for i < len(word) {
		j := i + 1
		cnt := 1
		for j < len(word) && word[i] == word[j] {
			cnt++
			if cnt == 9 {
				comm = append(comm, rune('0'+cnt))
				comm = append(comm, rune(word[i]))
				cnt = 0
			}
			j++
		}
		if cnt != 0 {
			comm = append(comm, rune('0'+cnt))
			comm = append(comm, rune(word[i]))
		}

		i = j
	}

	return string(comm)
}
